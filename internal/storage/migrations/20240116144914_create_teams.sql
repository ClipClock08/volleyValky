-- +goose Up
-- +goose StatementBegin
CREATE TABLE Teams
(
    TeamID       SERIAL PRIMARY KEY,
    TeamName     TEXT,
    TeamLocation TEXT,
    CoachName    TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Teams;
-- +goose StatementEnd
