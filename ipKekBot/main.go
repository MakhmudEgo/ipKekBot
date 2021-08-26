package main

import (
	"encoding/json"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"ipKekBot/connectDB"
	"ipKekBot/users"
	"net/http"
	"time"
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
	db.DB()
	bot, updates := Init()
	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			user := &users.Users{Id: update.Message.From.ID}
			prevMsg := db.First(user, "id = ?", user.Id)
			if prevMsg.Error != nil {
				db.Create(user)
			}
			fmt.Println(prevMsg)
			var reply string
			var msg tg.MessageConfig
			if update.Message.IsCommand() {
				switch update.Message.Text {
				case "/start":
					{
						reply = fmt.Sprintf("Yo @%s! I'm ipKekbot.\n"+
							"I can help you get information about the ip and store it in the database.\n\n"+
							"You can control me by sending these commands:\n\n"+
							"/checkip - get information about the ip\n"+
							"/historycheck - get history checks",
							update.Message.From.UserName)
						msg = tg.NewMessage(update.Message.Chat.ID, reply)
						if user.Adm == true {
							msg.ReplyMarkup = users.CreatorBtn
						} else {
							msg.ReplyMarkup = users.UserBtn
						}
						msg.ReplyToMessageID = update.Message.MessageID
					}
				case "/checkip":
					{
						reply = "Enter the ip address.\nExample: 213.180.193.1"
						msg = tg.NewMessage(update.Message.Chat.ID, reply)
						msg.ReplyMarkup = users.CreatorBtn
						msg.ReplyToMessageID = update.Message.MessageID
						tst := db.Find(&users.Users{Id: update.Message.From.ID}).Update("prev_msg", update.Message.Text)
						fmt.Println(tst)
					}
				default:

				}
			} else {

				switch user.PrevMsg {
				case "/checkip":
					msg = handleCheckIp(db, update.Message, &update.Message.Text)
					if user.Adm {
						msg.ReplyMarkup = users.CreatorBtn
					} else {
						msg.ReplyMarkup = users.UserBtn
					}
					user.PrevMsg = ""
					db.Save(user)
				}
			}

			mess, err := bot.Send(msg)
			if err != nil {
				log.Fatal(err, mess)
			}
		}
	}
}

func handleCheckIp(db *gorm.DB, message *tg.Message, query *string) tg.MessageConfig {
	var msg tg.MessageConfig
	var reply string
	d := &connectDB.Ips{Query: *query}
	res := db.First(d, "query = ?", *query)
	if res.Error != nil || d.Query != *query {
		kek, err := http.Get("http://ip-api.com/json/" + *query)
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(kek.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer kek.Body.Close()
		if errJ := json.Unmarshal(b, &d); errJ != nil {
			log.Fatal(errJ)
		}
		d.ID = 0
		res = db.Create(&d)
		if res.Error != nil {
			log.Fatal(res.Error)
		}
		if d.Status == "success" {
			reply = generateRespMsgText(d)
		} else {
			reply = fmt.Sprintf("Bad request!\nQuery Ip: %s", d.Query)
		}
	} else {
		reply = generateRespMsgText(d)
	}
	msg = tg.NewMessage(int64(message.From.ID), reply)
	msg.ReplyToMessageID = message.MessageID
	return msg
}

func generateRespMsgText(d *connectDB.Ips) string {
	return fmt.Sprintf(`Query Ip: %s
Region: %s
RegionName: %s
countryCode: %s
Country: %s
City: %s
Zip: %s
Lat: %f
Lon: %f
Timezone: %s
Isp: %s
Org: %s
As: %s`, d.Query, d.Region, d.RegionName, d.CountryCode, d.Country,
		d.City, d.Zip, d.Lat, d.Lon, d.Timezone, d.Isp, d.Org, d.As)
}
