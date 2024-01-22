-- +goose Up
-- +goose StatementBegin
CREATE TABLE MatchSets
(
    MatchSetID SERIAL PRIMARY KEY,
    MatchID    INT,
    SetNumber  INT,
    Team1Score INT,
    Team2Score INT,
    FOREIGN KEY (MatchID) REFERENCES Matches (MatchID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS MatchSets;
-- +goose StatementEnd
