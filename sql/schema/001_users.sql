-- +goose up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name varchar(255) UNIQUE NOT NULL
);

-- +goose down
DROP TABLE users;