-- +goose Up
-- +goose StatementBegin

ALTER TABLE users ADD COLUMN IF NOT EXISTS notify_interval BIGINT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS notified_at TIMESTAMP DEFAULT now();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE users DROP COLUMN IF EXISTS notify_interval;
ALTER TABLE users DROP COLUMN IF EXISTS notified_at;

-- +goose StatementEnd
