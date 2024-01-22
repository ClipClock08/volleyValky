package bot

import (
	"context"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"volleyvalkybot/internal/botkit"
	"volleyvalkybot/internal/botkit/markup"
	"volleyvalkybot/internal/model"
)

type SeasonProvider interface {
	SeasonByID(ctx context.Context, id int64) (*model.Season, error)
}

func ViewCmdGetSeason(provider SeasonProvider) botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		idStr := update.Message.CommandArguments()

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		season, err := provider.SeasonByID(ctx, id)
		if err != nil {
			return err
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, formatSeason(*season))
		reply.ParseMode = parseModeMarkdownV2

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}

func formatSeason(source model.Season) string {
	return fmt.Sprintf(
		"🌐 ID: `%d`\n:*%s*\n %s\nDescr: %s",
		source.SeasonID,
		source.StartDate,
		source.EndDate,
		markup.EscapeForMarkdown(source.Description),
	)
}
