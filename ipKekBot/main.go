package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"ipKekBot/Command"
	"ipKekBot/Handler"
	"ipKekBot/connectDB"
	"ipKekBot/users"
	"time"
)

func Init() (*tg.BotAPI, tg.UpdatesChannel) {
	token := "1971227789:AAEwskoVtcc3ap4qlXYMyqX0Jx3BJLv0G7g" // BAD PRACTICES

	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Fatal(err.Error())
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tg.NewUpdate(0)
	u.Timeout = int(time.Hour)
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bot, updates
}

func main() {
	db := connectDB.Connect()
	bot, updates := Init()
	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}

			// init || create user
			user := &users.Users{Id: update.Message.From.ID, Username: update.Message.From.UserName}
			prevMsg := db.First(user, "id = ?", user.Id)
			if prevMsg.Error != nil {
				tx := db.Create(user)
				if tx.Error != nil {
					log.Fatal(tx.Error.Error())
				}
			}

			var msg tg.MessageConfig
			if update.Message.IsCommand() {
				msg = Command.Check(user, db, update.Message, bot)
				if msg.ChatID == -1 {
					continue
				}
			} else {
				msg = Handler.Execute(user, db, update.Message)
			}

			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
