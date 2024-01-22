package model

import (
	"time"
)

type Item struct {
	Title      string
	Categories []string
	Link       string
	Date       time.Time
	Summary    string
	SourceName string
}

type Source struct {
	ID        int64
	Name      string
	FeedURL   string
	Priority  int
	CreatedAt time.Time
}

type Article struct {
	ID          int64
	SourceID    int64
	Title       string
	Link        string
	Summary     string
	PublishedAt time.Time
	PostedAt    time.Time
	CreatedAt   time.Time
}

type Season struct {
	SeasonID    int64
	StartDate   time.Time
	EndDate     time.Time
	Description string
}

type Player struct {
	ID       int64
	TeamID   int64
	Name     string
	Lastname string
	Surname  string
	Born     time.Time
	Position string
	Number   uint8
	Photo    string
}

type Coach struct {
	ID       int64
	Name     string
	Lastname string
	Surname  string
}

type Team struct {
	ID      int64
	Name    string
	Since   time.Time
	CoachID int64
	Players []Player
}

type Match struct {
}

type MatchResults struct {
}

type Result struct {
	Season  int64
	Teams   []Team
	Matches []Match
	Results []MatchResults
}
