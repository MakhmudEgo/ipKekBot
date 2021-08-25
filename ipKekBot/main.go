package main

import (
	"encoding/json"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Data struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

var numericKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonURL("1.com", "http://x.com"),
		tg.NewInlineKeyboardButtonData("hi", "/start"),
		tg.NewInlineKeyboardButtonData("3", "3"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("4", "4"),
		tg.NewInlineKeyboardButtonData("5", "5"),
		tg.NewInlineKeyboardButtonData("6", "6"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonSwitch("10", "10"),
	),
)


var myBtn = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/start"),
		tg.NewKeyboardButton("2"),
		tg.NewKeyboardButton("3"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("4"),
		tg.NewKeyboardButton("5"),
		tg.NewKeyboardButton("6"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButtonContact("7"),
		tg.NewKeyboardButtonLocation("8"),
		tg.NewKeyboardButtonContact("9"),
	),
)

func main() {
	token := "1971227789:AAEwskoVtcc3ap4qlXYMyqX0Jx3BJLv0G7g"
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	ipApi := "http://ip-api.com/json/"
	kek, _ := http.Get(ipApi + "xxxx")

	b, err := ioutil.ReadAll(kek.Body)
	var d Data

	if errJ := json.Unmarshal(b, &d); errJ != nil {
		log.Fatal(errJ)
	}
	fmt.Println(d)
	bot.Debug = true


	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tg.NewUpdate(0)
	u.Timeout = int(time.Hour)

	//msg := tg.NewMessage(234899515, "go go go")
	//bot.Send(msg)

	updates, err := bot.GetUpdatesChan(u)
	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				if update.ChannelPost != nil {

					resp, _ := bot.SetChatTitle(tg.SetChatTitleConfig{Title: "helloD world!", ChatID: update.ChannelPost.Chat.ID})
					bot.SetChatDescription(tg.SetChatDescriptionConfig{ChatID: update.ChannelPost.Chat.ID, Description: "kek lol"})
					admins, _ := bot.GetChatAdministrators(tg.ChatConfig{ChatID: update.ChannelPost.Chat.ID})
					fmt.Println(admins)

					fmt.Println("reeeeeeeesp", resp)
				}
				continue
			}
			if update.Message.IsCommand() && update.Message.Text == "/start" {
				reply := fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо.`,
					update.Message.From.UserName)
				msg := tg.NewMessage(update.Message.Chat.ID, reply)

				msg.ReplyMarkup = 	myBtn

				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)

				continue
			}
			UserName := update.Message.From.UserName
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			var reply string
			reply = fmt.Sprintf(`@%s, ну что!?! Веди себя хорошо.`,
					update.Message.From.UserName)

			msg := tg.NewMessage(ChatID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = numericKeyboard
			fmt.Println(msg.ChannelUsername)

			bot.Send(msg)
		}
	}
}
