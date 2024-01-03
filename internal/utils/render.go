package utils

import (
	"dollar-bot/internal/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RenderCotization(resp []types.MoneyDetail, flag string, msg *tgbotapi.MessageConfig) {

	msgText := ""
	for _, item := range resp {
		dateTime, err := FormatDate(item.UpdateAt)

		if err != nil {
			_ = fmt.Errorf(`Error format updatedAt: %s\n`, err.Error())
		}

		msgText += fmt.Sprintf(
			`%s *%s*:
						 - *Compra:* $%.2f
						 - *Venta:* $%.2f
						 - *Fecha Actualización:* %s
`,
			flag, item.Name, item.Buy, item.Sell, dateTime,
		)
		msgText += fmt.Sprintf("\n")
	}
	msg.Text = msgText
	msg.ParseMode = "Markdown"
}
