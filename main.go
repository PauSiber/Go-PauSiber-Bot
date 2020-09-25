package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	t "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

// commands keeps command-message peers.
// Key is command, value is message.
var commands map[string]string

// init loads command and env files.
func init() {
	// loads commands to struct from json
	// that is located at files/commands.json
	f, err := os.Open("./files/commands.json")
	if err != nil {
		log.Fatalf("Error occur while opening the commands json file.\nERROR: %s\n", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Error occur while reading the file.")
	}

	err = json.Unmarshal(b, &commands)
	if err != nil {
		log.Fatalf("Error occur while unmarshalling the json value to struct.\nERROR: %s\n", err)
	}

	// The error message is ignored
	// to be able to use system's environments inside of env file.
	_ = godotenv.Load()
}

func main() {
	bot, err := t.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatalln("There is an issue with telegram token.")
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := t.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := t.NewMessage(update.Message.Chat.ID, "")
		if !update.Message.IsCommand() {
			msg.Text = "You didn't even specify a command."
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Sending message is failed.\nERROR: %s\n", err)
			}
			continue
		}

		if message, ok := commands[update.Message.Command()]; ok {
			msg.Text = message // If the command is valid, set the message.
		} else {
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Sending message is failed.\nERROR: %s\n", err)
		}
	}
}
