-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS cases(
    id SERIAL PRIMARY KEY,
    filename VARCHAR,
    question VARCHAR NOT NULL,
    correct_answer INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS  answers(
    id SERIAL PRIMARY KEY,
    case_id INT REFERENCES cases(id),
    answer VARCHAR NOT NULL,
    number INT NOT NULL
);


CREATE TABLE  IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    nickname VARCHAR,
    created_at TIMESTAMP DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE answers;

DROP TABLE cases;

DROP TABLE users;

-- +goose StatementEnd
