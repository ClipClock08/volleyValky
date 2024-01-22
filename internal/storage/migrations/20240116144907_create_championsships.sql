-- +goose Up
-- +goose StatementBegin
CREATE TABLE Championships
(
    ChampionshipID       SERIAL PRIMARY KEY,
    SeasonID             INT,
    ChampionshipName     TEXT,
    ChampionshipLocation TEXT,
    StartDate            DATE,
    EndDate              DATE,
    FOREIGN KEY (SeasonID) REFERENCES Seasons (SeasonID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Championships;
-- +goose StatementEnd
