package bot

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"

	"volleyvalkybot/internal/botkit"
	"volleyvalkybot/internal/model"
)

type SeasonLister interface {
	Seasons(ctx context.Context) ([]model.Season, error)
}

func ViewCmdListSeason(lister SeasonLister) botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		season, err := lister.Seasons(ctx)
		if err != nil {
			return err
		}

		var (
			seasonInfos = lo.Map(season, func(season model.Season, _ int) string { return formatSeason(season) })
			msgText     = fmt.Sprintf(
				"Список сезонів \\(всього %d\\):\n\n%s",
				len(season),
				strings.Join(seasonInfos, "\n\n"),
			)
		)

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		reply.ParseMode = parseModeMarkdownV2

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}
