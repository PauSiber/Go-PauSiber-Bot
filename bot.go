package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"os"
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

	jsonFile, err := os.Open("commands.json")
	checkErr(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonDATA JSONData
	json.Unmarshal(byteValue, &jsonDATA)

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
		// https://apps.timwhitlock.info/emoji/tables/unicode

		for _, command := range jsonDATA.Commands {
			// TODO we need to use map, instead of using a loop.
			if update.Message.Command() == command.Botcommand {
				msg.Text = command.Botmessage
				break
			}
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("Mesaj gönderilemedi", err)
		}
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
}
