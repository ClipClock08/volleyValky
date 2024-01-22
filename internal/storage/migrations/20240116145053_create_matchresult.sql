-- +goose Up
-- +goose StatementBegin
CREATE TABLE MatchResult
(
    ResultID      SERIAL PRIMARY KEY,
    MatchID       INT,
    WinningTeamID INT,
    LosingTeamID  INT,
    FOREIGN KEY (MatchID) REFERENCES Matches (MatchID),
    FOREIGN KEY (WinningTeamID) REFERENCES Teams (TeamID),
    FOREIGN KEY (LosingTeamID) REFERENCES Teams (TeamID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS MatchResult;
-- +goose StatementEnd
