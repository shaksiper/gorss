-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE feed_follow (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feed(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE feed_follow;
-- +goose StatementEnd
