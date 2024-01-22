-- +goose Up
-- +goose StatementBegin
CREATE TABLE TeamSeason
(
    TeamSeasonID SERIAL PRIMARY KEY,
    SeasonID     INT,
    TeamID       INT,
    FOREIGN KEY (SeasonID) REFERENCES Seasons (SeasonID),
    FOREIGN KEY (TeamID) REFERENCES Teams (TeamID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS TeamSeason;
-- +goose StatementEnd
