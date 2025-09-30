package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

type (
	InfoRequest struct {
		Text string
	}

	llmResponse struct {
		Question string `json:"question"`
		Answers  []struct {
			Number int    `json:"number"`
			Text   string `json:"text"`
		} `json:"answers"`
		CorrectAnswer struct {
			Number int `json:"number"`
		} `json:"correct_answer"`
	}

	InfoResponse struct {
		Question      string
		Answers       []models.Answer
		CorrectAnswer int
	}
)

const (
	getInfoPrompt = `
Ты — эксперт по анализу экзаменационных и тестовых вопросов.  
Тебе дают неструктурированный текст вопроса с вариантами ответов.  
Твоя задача:  
1. Найти сам текст вопроса.  
2. Найти все варианты ответов.  
3. Определить правильный ответ (После ответа будет указан символ '*', что указывает на верный ответ).  
4. Вернуть результат строго в JSON-формате.  

Формат ответа:  
{
  "question": "<строка с формулировкой вопроса>",
  "answers": [
    {"number": 1, "text": "<вариант ответа>"},
    {"number": 2, "text": "<вариант ответа>"},
    {"number": 3, "text": "<вариант ответа>"}
  ],
  "correct_answer": {
    "number": 2
  }
}

Правила:  
- Возвращай только JSON, без лишних комментариев.
- После правильного ответа указан символ '*'.
- Не сокращай и не изменяй формулировки текста вопроса и ответов, перепиши их полностью.
- В случае с несколькими языками, выбирай русский вариант ответа.

Текст из которого следует достать информацию:
%s
`
)

func (c *Client) GetInfo(req InfoRequest) (*InfoResponse, error) {
	var out Out

	err := c.do(context.Background(), In{
		Model: c.gptModel,
		Messages: []Message{
			{
				Role:    "user",
				Content: fmt.Sprintf(getInfoPrompt, req.Text),
			},
		},
	},
		&out,
	)
	if err != nil {
		return nil, fmt.Errorf("get info: %w", err)
	}

	var llmResp llmResponse
	if err := json.Unmarshal([]byte(out.Choices[0].Message.Content), &llmResp); err != nil {
		return nil, fmt.Errorf("get info: %w", err)
	}

	var answers []models.Answer
	for _, a := range llmResp.Answers {
		answers = append(answers, models.Answer{
			Answer: a.Text,
			Number: a.Number,
		})
	}

	return &InfoResponse{
		Question:      llmResp.Question,
		Answers:       answers,
		CorrectAnswer: llmResp.CorrectAnswer.Number,
	}, nil
}
