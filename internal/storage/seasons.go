package storage

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"volleyvalkybot/internal/model"
)

type SeasonsPostgresStorage struct {
	db *sqlx.DB
}

func NewSeasonStorage(db *sqlx.DB) *SeasonsPostgresStorage {
	return &SeasonsPostgresStorage{db: db}
}

func (s *SeasonsPostgresStorage) Seasons(ctx context.Context) ([]model.Season, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()

	var seasons []dbSeason
	if err := conn.SelectContext(ctx, &seasons, `SELECT * FROM seasons`); err != nil {
		return nil, err
	}

	return lo.Map(seasons, func(season dbSeason, _ int) model.Season { return model.Season(season) }), nil
}

func (s *SeasonsPostgresStorage) SeasonByID(ctx context.Context, id int64) (*model.Season, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()

	var season dbSeason
	if err := conn.GetContext(ctx, &season, `SELECT * FROM seasons WHERE seasonid = $1`, id); err != nil {
		return nil, err
	}

	return (*model.Season)(&season), nil
}

func (s *SeasonsPostgresStorage) Add(ctx context.Context, season model.Season) (int64, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer func() { _ = conn.Close() }()

	var id int64

	row := conn.QueryRowxContext(
		ctx,
		`INSERT INTO seasons (startdate, enddate, description)
					VALUES ($1, $2, $3) RETURNING seasonid;`,
		season.StartDate, season.EndDate, season.Description,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SeasonsPostgresStorage) SetDescription(ctx context.Context, id int64, priority string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	_, err = conn.ExecContext(ctx, `UPDATE seasons SET description = $1 WHERE seasonid = $2`, priority, id)

	return err
}

func (s *SeasonsPostgresStorage) Delete(ctx context.Context, id int64) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	if _, err := conn.ExecContext(ctx, `DELETE FROM seasons WHERE seasonid = $1`, id); err != nil {
		return err
	}

	return nil
}

type dbSeason struct {
	SeasonID    int64     `db:"seasonid"`
	StartDate   time.Time `db:"startdate"`
	EndDate     time.Time `db:"enddate"`
	Description string    `db:"description"`
}
