-- +goose Up
-- +goose StatementBegin
CREATE TABLE Matches
(
    MatchID        SERIAL PRIMARY KEY,
    ChampionshipID INT,
    Team1ID        INT,
    Team2ID        INT,
    MatchDate      DATE,
    FOREIGN KEY (ChampionshipID) REFERENCES Championships (ChampionshipID),
    FOREIGN KEY (Team1ID) REFERENCES Teams (TeamID),
    FOREIGN KEY (Team2ID) REFERENCES Teams (TeamID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Matches;
-- +goose StatementEnd
