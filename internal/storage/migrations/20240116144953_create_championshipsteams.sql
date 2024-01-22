-- +goose Up
-- +goose StatementBegin
CREATE TABLE ChampionshipsTeams
(
    ChampionshipTeamID SERIAL PRIMARY KEY,
    ChampionshipID     INT,
    TeamID             INT,
    FOREIGN KEY (ChampionshipID) REFERENCES Championships (ChampionshipID),
    FOREIGN KEY (TeamID) REFERENCES Teams (TeamID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Championships;
-- +goose StatementEnd
