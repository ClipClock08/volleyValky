-- +goose Up
-- +goose StatementBegin
CREATE TABLE Seasons
(
    SeasonID    SERIAL PRIMARY KEY,
    StartDate   DATE NOT NULL,
    EndDate     DATE NOT NULL,
    Description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Seasons;
-- +goose StatementEnd
