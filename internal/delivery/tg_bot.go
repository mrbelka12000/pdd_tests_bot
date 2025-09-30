package delivery

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
	"github.com/mrbelka12000/pdd_tests_bot/internal/usecase"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/config"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/pointer"
)

type (
	Handler struct {
		uc  *usecase.UseCase
		log *slog.Logger

		bot *tgbotapi.BotAPI
		ch  tgbotapi.UpdatesChannel

		cache cache
	}

	cache interface {
		Set(key string, value interface{}, dur time.Duration) error
		Get(key string) (string, bool)
		GetInt64(key string) (int64, bool)
		GetInt(key string) (int, bool)
		Delete(key string)
	}
)

func Start(cfg config.Config, uc *usecase.UseCase, log *slog.Logger, cache cache) (*Handler, error) {

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, fmt.Errorf("new bot: %w", err)
	}

	uCfg := tgbotapi.NewUpdate(0)
	uCfg.Timeout = 60

	h := Handler{
		uc:    uc,
		log:   log,
		bot:   bot,
		ch:    bot.GetUpdatesChan(uCfg),
		cache: cache,
	}
	go h.sendQuizzes()

	h.handleUpdate()

	return &h, nil
}

func (h *Handler) handleUpdate() {
	for update := range h.ch {

		if update.CallbackQuery != nil {
			h.handleCallbacks(update.CallbackQuery)
			continue
		}

		if update.Message == nil {
			continue
		}

		msg := update.Message

		switch msg.Command() {
		case "start":

			msgToSend := tgbotapi.NewMessage(msg.Chat.ID, "Как часто хотите получать тесты ?")
			msgToSend.ReplyMarkup = h.getNotifyInterval()

			sentMsg, err := h.bot.Send(msgToSend)
			if err != nil {
				h.handleSendMessageError(h.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ошибка")))
				continue
			}

			h.cache.Set(fmt.Sprint(msg.Chat.ID), sentMsg.MessageID, 5*time.Minute)
		}
	}
}

func (h *Handler) handleCallbacks(cb *tgbotapi.CallbackQuery) {

	h.log.Info("handleCallbacks")

	cbData, err := unmarshalCallbackData(cb.Data)
	if err != nil {
		h.log.With("error", err).Error("unmarshal callback data")
		return
	}

	switch {
	case cbData.Inter != nil:
		interval, err := time.ParseDuration(cbData.Inter.Val)
		if err != nil {
			h.log.With("error", err).Error("parse interval")
			h.handleSendMessageError(h.bot.Send(tgbotapi.NewMessage(cb.Message.Chat.ID, "Ошибка")))
			return
		}

		origUser, _ := h.uc.GetUserByChatID(cb.Message.Chat.ID)

		user := models.User{
			ID:             origUser.ID,
			ChatID:         cb.Message.Chat.ID,
			Nickname:       cb.Message.From.UserName,
			CreatedAt:      time.Now().UTC(),
			NotifyInterval: interval,
			NotifiedAt:     time.Now().UTC(),
		}

		if err := h.uc.CreateUser(user); err != nil {
			h.log.With("error", err).Error("create user")
			h.handleSendMessageError(h.bot.Send(tgbotapi.NewMessage(cb.Message.Chat.ID, "Ошибка")))
			return
		}

		messageID, _ := h.cache.GetInt(fmt.Sprint(cb.Message.Chat.ID))
		msgToDelete := tgbotapi.NewDeleteMessage(cb.Message.Chat.ID, messageID)

		h.cache.Delete(fmt.Sprint(cb.Message.Chat.ID))
		h.handleSendMessageError(h.bot.Send(msgToDelete))

	case cbData.A != nil:
		cs, err := h.uc.GetCase(cbData.A.CaseID)
		if err != nil {
			h.log.With("error", err).Error("get case")
			return
		}

		if cbData.A.AnswerNum != cs.CorrectAnswer {
			h.handleSendMessageError(h.bot.Send(tgbotapi.NewCallback(cb.ID, "Не правильный ответ")))
		} else {
			h.handleSendMessageError(h.bot.Send(tgbotapi.NewCallback(cb.ID, "Правильный ответ")))
		}
	}
}

func (h *Handler) handleSendMessageError(_ tgbotapi.Message, err error) {
	if err != nil {
		h.log.With("error", err).Error("send message")
	}
}

func (h *Handler) sendQuizzes() {

	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		users, err := h.uc.GetAllUsers()
		if err != nil {
			h.log.With("error", err).Error("get all users")
			continue
		}

		for _, user := range users {
			if time.Now().Before(user.NotifiedAt) {
				continue
			}

			cs, err := h.uc.GetRandomCase()
			if err != nil {
				h.log.With("error", err).Error("get random case")
				continue
			}

			if pointer.Value(cs.Filename) == "" {
				continue
			}

			var (
				buttons  [][]tgbotapi.InlineKeyboardButton
				button   = make([]tgbotapi.InlineKeyboardButton, 0, 2)
				response strings.Builder
				numbers  = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "1️⃣0️⃣"}
			)

			for i, ans := range cs.Answers {
				response.WriteString(fmt.Sprintf("%s: %s\n", numbers[i], ans.Answer))
				button = append(button, tgbotapi.InlineKeyboardButton{
					Text: fmt.Sprintf("Choose %s", numbers[i]),
					CallbackData: pointer.Of(marshalCallbackData(CallbackData{
						A: &Answer{
							AnswerNum: ans.Number,
							CaseID:    ans.CaseID,
						},
					})),
				})

				if len(button) == 2 {
					buttons = append(buttons, button)
					button = make([]tgbotapi.InlineKeyboardButton, 0, 2)
				}
			}

			if len(button) > 0 {
				buttons = append(buttons, button)
			}

			filename, err := h.uc.DownloadFile(*cs.Filename)
			if err != nil {
				h.log.With("error", err).Error("download file")
			}

			if filename != "" {
				photo := tgbotapi.NewPhoto(user.ChatID, tgbotapi.FilePath(filename))
				photo.Caption = cs.Question
				h.handleSendMessageError(h.bot.Send(photo))
				if err := os.Remove(filename); err != nil {
					h.log.With("error", err).Error("remove file")
				}
			} else {
				msgToSend := tgbotapi.NewMessage(user.ChatID, cs.Question)
				h.handleSendMessageError(h.bot.Send(msgToSend))
			}

			messageToSend := tgbotapi.NewMessage(user.ChatID, response.String())
			messageToSend.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons...)
			h.handleSendMessageError(h.bot.Send(messageToSend))

			h.log.Info("send quizzes")

			if err := h.uc.UpdateUser(models.User{
				ID:         user.ID,
				NotifiedAt: time.Now().UTC().Add(user.NotifyInterval),
			}); err != nil {
				h.log.With("error", err).Error("update user")
			}
		}
	}
}

func marshalCallbackData(cb CallbackData) string {
	body, _ := json.Marshal(cb)
	return string(body)
}

func unmarshalCallbackData(data string) (cb CallbackData, err error) {
	err = json.Unmarshal([]byte(data), &cb)
	return
}
