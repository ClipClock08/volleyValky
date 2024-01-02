package model

import "time"

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
	ID      int64
	Name    string
	Teams   []Team
	ScoreID int64
}

type Team struct {
	ID      int64
	Name    string
	Players []Player
}

type Player struct {
	ID        int64
	Name      string
	Surname   string
	Lastname  string
	Photo     string
	BirthDate time.Time
	Role      Role
	Number    int
}

type Role struct {
	ID   int64
	Role string
}

type Score struct {
	ID         int64
	SeasonID   int64
	Team1ID    int64
	Team2ID    int64
	ScoreTeam1 int8
	ScoreTeam2 int8
}
