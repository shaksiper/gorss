-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE feed ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE feed DROP COLUMN last_fetched_at;
-- +goose StatementEnd
