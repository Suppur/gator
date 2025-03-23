-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    feed_id uuid REFERENCES feeds(id) ON DELETE CASCADE,
    CONSTRAINT user_feed UNIQUE (user_id, feed_id)
);


-- +goose Down
DROP TABLE feed_follows;