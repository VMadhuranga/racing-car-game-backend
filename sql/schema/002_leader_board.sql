-- +goose Up
CREATE TABLE leader_board (
    id UUID PRIMARY KEY,
    best_time TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE leader_board;