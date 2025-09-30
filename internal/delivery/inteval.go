package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/mrbelka12000/pdd_tests_bot/pkg/pointer"
)

func (h *Handler) getNotifyInterval() tgbotapi.InlineKeyboardMarkup {

	var (
		buttons [][]tgbotapi.InlineKeyboardButton
	)

	buttons = append(buttons, []tgbotapi.InlineKeyboardButton{
		{
			Text: "1 Минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "1m"},
			})),
		},
		{
			Text: "2 Минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "2m"},
			})),
		},
		{
			Text: "3 Минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "3m"},
			})),
		},
	})

	buttons = append(buttons, []tgbotapi.InlineKeyboardButton{
		{
			Text: "5 Минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "5m"},
			})),
		},
		{
			Text: "10 Минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "10m"},
			})),
		},
		{
			Text: "15 Минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "15m"},
			})),
		},
	})

	buttons = append(buttons, []tgbotapi.InlineKeyboardButton{
		{
			Text: "20 минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "20m"},
			})),
		},
		{
			Text: "25 минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "25m"},
			})),
		},
		{
			Text: "30 минут",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "30m"},
			})),
		},
	})

	buttons = append(buttons, []tgbotapi.InlineKeyboardButton{
		{
			Text: "1 Час",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "1h"},
			})),
		},
		{
			Text: "2 часа",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "2h"},
			})),
		},
		{
			Text: "3 часа",
			CallbackData: pointer.Of(marshalCallbackData(CallbackData{
				Inter: &Interval{Val: "3h"},
			})),
		},
	})

	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
