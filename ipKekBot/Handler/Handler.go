package Handler

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

func Execute(user *users.Users, db *gorm.DB, upd *tg.Message, bot *tg.BotAPI) tg.MessageConfig {
	var msg tg.MessageConfig

	switch user.PrevMsg {
	case "/checkip":
		msg = respCheckIp(db, upd)
		user.PrevMsg = ""
		db.Save(user)
	case "/sendmessage":
		sendmessage(db, upd, bot)
		user.PrevMsg = ""
		db.Save(user)
		return tg.NewMessage(-1, "")
	case "/addnewadmin":
		msg = addNewAdmin(db, upd)
		user.PrevMsg = ""
		db.Save(user)
	default:
		msg = otherMsg(user.Role, upd)
	}
	msg.ReplyToMessageID = upd.MessageID
	return msg
}

func addNewAdmin(db *gorm.DB, upd *tg.Message) tg.MessageConfig {
	newAdm := &users.Users{Username: upd.Text}
	var msg tg.MessageConfig
	res := db.First(newAdm, "username = ?", upd.Text)
	if res.Error != nil {
		msg = tg.NewMessage(int64(upd.From.ID), "Not found "+upd.Text)
	} else {
		msg = tg.NewMessage(int64(upd.From.ID), "Added")
		newAdm.Role = 1
		db.Save(newAdm)
	}
	msg.ReplyToMessageID = upd.MessageID
	return msg
}

func sendmessage(db *gorm.DB, upd *tg.Message, bot *tg.BotAPI) {
	var usrs []users.Users
	db.Select("id").Find(&usrs)
	reply := upd.Text
	for _, usr := range usrs {
		if usr.Id == upd.From.ID {
			continue
		}
		msg := tg.NewMessage(int64(usr.Id), reply)
		_, err := bot.Send(msg)
		if err != nil && err.Error() != "Forbidden: bot was blocked by the user" {
			log.Fatal(err.Error())
		}
	}
}

func otherMsg(role int, upd *tg.Message) tg.MessageConfig {
	msg := tg.NewMessage(int64(upd.From.ID),
		"Я не болталка - я Бот деловой!\nВас Бог знает сколько, а я всегда один.\nДелов и так хватает!\nПоэтому пишите исключительно по делу для которых я предназначен!\nСпасибо!")
	msg.ReplyMarkup = users.CheckRole(role)
	return msg

}

func respCheckIp(db *gorm.DB, message *tg.Message) tg.MessageConfig {
	var msg tg.MessageConfig
	var reply string
	d := &connectDB.Ips{Query: message.Text}
	res := db.First(d, "query = ?", message.Text)
	if res.Error != nil || d.Query != message.Text {
		sendRequest(d, &message.Text)
		res = db.Create(d)
		if res.Error != nil {
			log.Fatal(res.Error.Error())
		}
		reply = GenerateRespCheckIp(d)
	} else {
		reply = GenerateRespCheckIp(d)
	}
	msg = tg.NewMessage(int64(message.From.ID), reply)

	tx := db.Create(&users.UserHistory{IpsId: d.ID, UserId: message.From.ID, Time: time.Now()})
	if tx.Error != nil {
		log.Fatal(tx.Error.Error())
	}

	return msg
}

func sendRequest(d *connectDB.Ips, query *string) {
	resp, err := http.Get("http://ip-api.com/json/" + *query)
	if err != nil {
		log.Fatal(err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := resp.Body.Close(); err != nil {
		log.Fatal(err.Error())
	}
	if err := json.Unmarshal(b, d); err != nil {
		log.Fatal(err.Error())
	}
}

func GenerateRespCheckIp(d *connectDB.Ips) string {
	if d.Status == "fail" {
		return fmt.Sprintf("Bad request!\nQuery Ip: %s", d.Query)
	}
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
