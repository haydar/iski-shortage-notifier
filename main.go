package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haydar/iski-incident-notifier/iski"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	//incidents := getAllIncidents();

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "aktifkesintiler":
			msg.ParseMode = "MarkdownV2"
			msg.Text = getActiveIncident()
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}

}

func getActiveIncident() string {
	incidents := iski.GetAllShortage()

	tmp := `
	*KESİNTİ LİSTESİ*
	`

	for _, data := range incidents.Data {

		tmp = tmp +
			"\n" +
			"\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_\\_" +
			"\n"+ 
			`*` + cases.Title(language.Turkish).String(data.IlceAdi) + `*` 

		for _, detail := range data.Detail {
			tmp = tmp +
				"\n \n" +
				`*Mahalle: *` + escapeDots(cases.Title(language.Turkish).String(detail.MahalleAdi)) +
				"\n" +
				`*Arıza Detayı: *` + escapeDots(cases.Title(language.Turkish).String(detail.ArizaNeviAciklamasi)) +
				"\n" +
				`*Başlama Tarihi: *` + escapeDots(cases.Title(language.Turkish).String(detail.BaslamaTarihi)) +
				"\n" +
				`*Tahmini Bitiş Tarihi: *` + escapeDots(cases.Title(language.Turkish).String(detail.TahminiBitisTarihi))
		}

	}

	return tmp
}

func escapeDots(input string) string {
	escaped := strings.ReplaceAll(input, ".", "\\.")
	return escaped
}
