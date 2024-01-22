-- +goose Up
-- +goose StatementBegin
CREATE TABLE Players
(
    PlayerID       SERIAL PRIMARY KEY,
    PlayerName     TEXT,
    PlayerPosition TEXT,
    BirthDate      DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Players;
-- +goose StatementEnd
