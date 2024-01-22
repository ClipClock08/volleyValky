package bot

import (
	"context"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"volleyvalkybot/internal/botkit"
	"volleyvalkybot/internal/model"
)

type SeasonStorage interface {
	Add(ctx context.Context, source model.Season) (int64, error)
}

func ViewCmdAddSeason(storage SeasonStorage) botkit.ViewFunc {
	type addSeasonArgs struct {
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		Description string    `json:"description"`
	}

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		args, err := botkit.ParseJSON[addSeasonArgs](update.Message.CommandArguments())
		if err != nil {
			return err
		}

		season := model.Season{
			StartDate:   args.StartDate,
			EndDate:     args.EndDate,
			Description: args.Description,
		}

		seasonID, err := storage.Add(ctx, season)
		if err != nil {
			// TODO: send error message
			return err
		}

		var (
			msgText = fmt.Sprintf(
				"Сезон доданий з ID: `%d`\\. Використовуйте цей ID для оновлення чи видалення сезону\\.",
				seasonID,
			)
			reply = tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		)

		reply.ParseMode = parseModeMarkdownV2

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}
