-- +goose Up
CREATE TABLE feeds (
    id uuid PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;