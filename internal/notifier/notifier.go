package notifier

import (
	"context"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"volleyvalkybot/internal/botkit/markup"
	"volleyvalkybot/internal/model"
)

type ResultsProvider interface {
	GetActiveSeason(ctx context.Context) (model.Season, error)
	GetTeams(ctx context.Context, season int64) ([]model.Team, error)
	GetMatches(ctx context.Context, season int64) ([]model.Match, error)
	GetMatchesSets(ctx context.Context, season int64) ([]model.MatchResults, error)
	MarkAsPosted(ctx context.Context, article model.Result) error
}

type Notifier struct {
	results          ResultsProvider
	bot              *tgbotapi.BotAPI
	sendInterval     time.Duration
	lookupTimeWindow time.Duration
	channelID        int64
}

func New(
	resultsProvider ResultsProvider,
	bot *tgbotapi.BotAPI,
	sendInterval time.Duration,
	lookupTimeWindow time.Duration,
	channelID int64,
) *Notifier {
	return &Notifier{
		results:          resultsProvider,
		bot:              bot,
		sendInterval:     sendInterval,
		lookupTimeWindow: lookupTimeWindow,
		channelID:        channelID,
	}
}

func (n *Notifier) Start(ctx context.Context) error {
	ticker := time.NewTicker(n.sendInterval)
	defer ticker.Stop()

	if err := n.SelectAndSendResults(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
			if err := n.SelectAndSendResults(ctx); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *Notifier) SelectAndSendResults(ctx context.Context) error {
	season, err := n.results.GetActiveSeason(ctx)
	if err != nil {
		return err
	}

	teams, err := n.results.GetTeams(ctx, season.SeasonID)
	if err != nil {
		return err
	}

	if len(teams) == 0 {
		return nil
	}

	matches, err := n.results.GetMatches(ctx, season.SeasonID)
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		return nil
	}

	matchesSets, err := n.results.GetMatchesSets(ctx, season.SeasonID)
	if err != nil {
		return err
	}

	if len(matchesSets) == 0 {
		return nil
	}

	result := prepareResult(season.SeasonID, teams, matches, matchesSets)

	if err := n.sendArticle(result); err != nil {
		return err
	}

	return n.results.MarkAsPosted(ctx, result)
}

func prepareResult(season int64, teams []model.Team, matches []model.Match, matchResults []model.MatchResults) model.Result {
	var result model.Result

	result = model.Result{
		Season:  season,
		Teams:   teams,
		Matches: matches,
		Results: matchResults,
	}

	return result
}

func (n *Notifier) sendArticle(result model.Result) error {
	const msgFormat = "*%s*%s\n\n%s"

	msg := tgbotapi.NewMessage(n.channelID, fmt.Sprintf(
		msgFormat,
		markup.EscapeForMarkdown(result.Teams[0].Name),
		markup.EscapeForMarkdown(result.Teams[0].Name),
		markup.EscapeForMarkdown(result.Teams[0].Name),
	))
	msg.ParseMode = "MarkdownV2"

	_, err := n.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
