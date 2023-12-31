package pkg

import (
	"dollar-bot/internal"
	"dollar-bot/internal/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
)

type Handler struct {
	Processor internal.Processor
}

func NewHandler() *Handler {
	return &Handler{
		Processor: internal.Processor{},
	}
}

func (h *Handler) HandleMessage() {
	bot, u := ConfigBot()

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyToMessageID = update.Message.MessageID

		switch update.Message.Text {
		case "/dolar":
			resp, err := h.Processor.GetCotizations(utils.Dollar)

			if err != nil {
				msg.Text = err.Error()
				continue
			}

			utils.RenderCotization(resp, "💵", &msg)
		case "/euro":
			resp, err := h.Processor.GetCotizations(utils.Euro)
			if err != nil {
				msg.Text = err.Error()
				continue
			}
			utils.RenderCotization(resp, "💶", &msg)
		case "/real":
			resp, err := h.Processor.GetCotizations(utils.Real)
			if err != nil {
				msg.Text = err.Error()
				continue
			}

			utils.RenderCotization(resp, "🇧🇷", &msg)
		case "/peso_ch":
			resp, err := h.Processor.GetCotizations(utils.PesoChilean)
			if err != nil {
				msg.Text = err.Error()
				continue
			}

			utils.RenderCotization(resp, "🇨🇱", &msg)
		case "/peso_uy":
			resp, err := h.Processor.GetCotizations(utils.PesoUruguayan)
			if err != nil {
				msg.Text = err.Error()
				continue
			}

			utils.RenderCotization(resp, "🇺🇾", &msg)
		default:
			msg.Text = utils.MessageDefault
		}

		bot.Send(msg)
	}
}

func ConfigBot() (*tgbotapi.BotAPI, tgbotapi.UpdateConfig) {
	fmt.Println("token: " + os.Getenv("TELEGRAM_TOKEN"))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	description := tgbotapi.SetChatTitleConfig{
		ChatID: 1,
		Title:  "Obtene las cotizaciones de las diferentes monedas 🇺🇸🇧🇷🇪🇺🇨🇱🇺🇾",
	}
	bot.SetChatTitle(description)
	if err != nil {
		log.Panic(err)
	}
	debug, _ := strconv.ParseBool(os.Getenv("MODE_DEBUG"))
	bot.Debug = debug

	log.Printf("Autorizado como %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot, u
}
