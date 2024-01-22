-- +goose Up
-- +goose StatementBegin
CREATE TABLE PlayerTeams
(
    PlayerTeamID SERIAL PRIMARY KEY,
    PlayerID     INT,
    TeamID       INT,
    SeasonID     INT,
    FOREIGN KEY (PlayerID) REFERENCES Players (PlayerID),
    FOREIGN KEY (TeamID) REFERENCES Teams (TeamID),
    FOREIGN KEY (SeasonID) REFERENCES Seasons (SeasonID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PlayerTeams;
-- +goose StatementEnd
