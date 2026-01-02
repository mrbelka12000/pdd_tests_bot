-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS messages(
  id SERIAL PRIMARY KEY,
  chat_id INT NOT NULL,
  telegram_message_id INT NOT NULL,
  case_id INT REFERENCES cases(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


-- +goose StatementEnd
