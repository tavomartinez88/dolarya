package utils

import (
	"dollar-bot/internal/types"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func RenderCotization(resp []types.MoneyDetail, flag string, msg *tgbotapi.MessageConfig) {

	msgText := ""
	for _, item := range resp {
		dateTimeSplited := strings.Split(item.UpdateAt, "T")
		time := strings.Split(dateTimeSplited[1], ":0")
		dateTime := dateTimeSplited[0] + " " + time[0]
		msgText += fmt.Sprintf(
			`
							%s *%s*:
							 - *Compra:* $%.2f
							 - *Venta:* $%.2f
							 - *Fecha Actualización:* %s
							`,
			flag, item.Name, item.Buy, item.Sell, dateTime,
		)
	}
	msg.Text = msgText
	msg.ParseMode = "Markdown"
}
