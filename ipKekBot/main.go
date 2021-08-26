package main

import (
	"encoding/json"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"ipKekBot/connectDB"
	"net/http"
	"time"
)

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

type UserHistory struct {
	IpsId  uint `gorm:"column:ips_id"`
	UserId uint `gorm:"column:user_id"`
}

func Init() (*tg.BotAPI, tg.UpdatesChannel) {
	token := "1971227789:AAEwskoVtcc3ap4qlXYMyqX0Jx3BJLv0G7g" // BAD PRACTICES

	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tg.NewUpdate(0)
	u.Timeout = int(time.Hour)
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	return bot, updates
}

func main() {
	db := connectDB.Connect()
	db.DB() // depr
	bot, updates := Init()

	// test service
	ipApi := "http://ip-api.com/json/"
	kek, _ := http.Get(ipApi + "195.133.239.83")
	b, _ := ioutil.ReadAll(kek.Body)
	defer kek.Body.Close()
	var d connectDB.Ips
	if errJ := json.Unmarshal(b, &d); errJ != nil {
		log.Fatal(errJ)
	}
	res := db.Create(&d) //res :=
	//tst:= db.First(&d)

	fmt.Println("====", res)
	//db.Create(&UserHistory{, 3})
	fmt.Println(d)

	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			var reply string
			var msg tg.MessageConfig
			switch update.Message.Text {
			case "/start":
				{
					reply = fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо.`,
						update.Message.From.UserName)
					msg = tg.NewMessage(update.Message.Chat.ID, reply)
					msg.ReplyMarkup = myBtn
					msg.ReplyToMessageID = update.Message.MessageID
				}
			case "/checkIp":
				{

				}
			default:

			}
			UserName := update.Message.From.UserName
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			reply = fmt.Sprintf(`@%s, ну что!?! Веди себя хорошо.`,
				update.Message.From.UserName)

			msg = tg.NewMessage(ChatID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = numericKeyboard
			fmt.Println(msg.ChannelUsername)

			bot.Send(msg)
		}
	}
}
