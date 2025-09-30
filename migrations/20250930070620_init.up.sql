-- +goose Up
-- +goose StatementBegin

CREATE TABLE cases(
    id SERIAL PRIMARY KEY,
    filename VARCHAR,
    question VARCHAR NOT NULL,
    correct_answer INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE answers(
    id SERIAL PRIMARY KEY,
    case_id INT REFERENCES cases(id),
    answer VARCHAR NOT NULL,
    number INT NOT NULL
);


CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    chat_id VARCHAR(50) NOT NULL,
    nickname VARCHAR,
    created_at TIMESTAMP DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE answers;

DROP TABLE cases;

-- +goose StatementEnd
