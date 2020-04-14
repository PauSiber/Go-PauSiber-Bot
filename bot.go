package main

import (
	"fmt"
    "log"
    _ "os"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    bot, err := tgbotapi.NewBotAPI("TOKEN")
    checkErr(err)

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)
	checkErr(err)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        if !update.Message.IsCommand() {
            continue
        }

        // Create a new MessageConfig. We don't have text yet, so we should leave it empty.
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

        // Extract the command from the Message.
        switch update.Message.Command() {
		// https://apps.timwhitlock.info/emoji/tables/unicode
        case "help":
			msg.Text =  "\xF0\x9F\xA6\x89 PauSiber Go Bot Komutlar \xF0\x9F\xA6\x89"
			msg.Text +=	"\n /sayhi \xF0\x9F\x91\x8B"
			msg.Text +=	"\n /status \xF0\x9F\x9A\x80"
			msg.Text += "\n /nice \xF0\x9F\x91\x8D"
			msg.Text += "\n /linux \xF0\x9F\x90\xA7"
			msg.Text += "\n /plan \xF0\x9F\x93\x91"
			msg.Text += "\n /kaynak \xF0\x9F\x93\x96"
        case "sayhi":
            msg.Text = "Hi :)"
        case "status":
			msg.Text = "I'm ok."
		case "google":
			msg.Text = "Google'dan araştırabilirsin dostum"
		case "nice":
			msg.Text = "Thank you bro :)"
		case "linux":
			msg.Text = "https://gnulinux.pausiber.xyz/"
		case "plan":
			msg.Text = "https://plan.pausiber.xyz/"
		case "kaynak":
			msg.Text = "https://kaynak.pausiber.xyz/"
		case "social":
			msg.Text =  "Web Sayfamız : https://pausiber.xyz/"
			msg.Text += "\n Twitter Hesabımız : https://twitter.com/siberpau"
			msg.Text += "\n GitHub Hesabımız : https://github.com/PauSiber"
			msg.Text += "\n Instagram Hesabımız : https://www.instagram.com/pausiber/"
		default:
            msg.Text = "I don't know that command"
        }

        if _, err := bot.Send(msg); err != nil {
            fmt.Printf("Mesaj gönderilemedi",err)
        }
    }
}

func checkErr(err error) {
	if err != nil {
			fmt.Println(err)
			//os.Exit(1)
	}
}