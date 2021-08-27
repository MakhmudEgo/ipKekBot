package Command

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ipKekBot/Handler"
	"ipKekBot/connectDB"
	"ipKekBot/users"
)

func Check(user *users.Users, db *gorm.DB, upd *tg.Message, bot *tg.BotAPI) tg.MessageConfig {
	var msg tg.MessageConfig

	switch upd.Text {
	case "/start":
		msg = start_(user, upd)
	case "/checkip":
		msg = checkIp_(db, upd)
	case "/historycheck":
		{
			msgs := historyCheck_(db, upd, user)
			for _, config := range msgs {
				_, err := bot.Send(config)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
			user.PrevMsg = ""
			db.Save(user)
		}
		return tg.NewMessage(-1, "")
	case "/sendmessage":
		if user.Role == users.UserR {
			msg = tg.NewMessage(int64(user.Id), "permission denied")
			msg.ReplyToMessageID = upd.MessageID
			return msg
		}
		msg = sendMessage_(db, upd)
		//::TODO duplicate addnewadmin, deleteadmin
	case "/addnewadmin":
		if user.Role == users.UserR {
			msg = tg.NewMessage(int64(user.Id), "permission denied")
			msg.ReplyToMessageID = upd.MessageID
			return msg
		}
		msg = tg.NewMessage(upd.Chat.ID, "enter the username")
		res := db.Find(&users.Users{Id: upd.From.ID}).Update("prev_msg", upd.Text)
		if res.Error != nil {
			log.Fatal(res.Error.Error())
		}
	case "/deleteadmin":
		if user.Role != users.CreatorR {
			msg = tg.NewMessage(int64(user.Id), "permission denied")
			msg.ReplyToMessageID = upd.MessageID
			return msg
		}
		msg = tg.NewMessage(upd.Chat.ID, "enter the username")
		res := db.Find(&users.Users{Id: upd.From.ID}).Update("prev_msg", upd.Text)
		if res.Error != nil {
			log.Fatal(res.Error.Error())
		}
	default:
		msg = unknown_(upd)
	}
	return msg
}

func sendMessage_(db *gorm.DB, upd *tg.Message) tg.MessageConfig {
	msg := tg.NewMessage(upd.Chat.ID, "enter a message")
	res := db.Find(&users.Users{Id: upd.From.ID}).Update("prev_msg", upd.Text)
	if res.Error != nil {
		log.Fatal(res.Error.Error())
	}
	return msg
}

func historyCheck_(db *gorm.DB, upd *tg.Message, user *users.Users) []tg.MessageConfig {
	var msgs []tg.MessageConfig
	var hist []users.UserHistory
	tx := db.Where("user_id = ?", upd.From.ID).Find(&hist)
	if tx.Error != nil {
		log.Fatal("historyCheck_\n" + tx.Error.Error())
	}
	myd := make(map[uint]connectDB.Ips)
	for _, item := range hist {
		var tmp connectDB.Ips
		if _, ok := myd[uint(item.UserId)]; !ok {
			db.Where("id = ?", item.IpsId).Find(&tmp)
			myd[tmp.ID] = tmp
		}
	}
	for idx, item := range hist {
		kek := myd[(item.IpsId)]
		reply := fmt.Sprintf("%d. %s\n\n%s",
			idx, item.Time.Format("Mon, 2 Jan 2006 15:04:05 MST"),
			Handler.GenerateRespCheckIp(&kek))
		msg := tg.NewMessage(int64(upd.From.ID), reply)
		msg.ReplyMarkup = users.CheckRole(user.Role)
		msg.ReplyToMessageID = upd.MessageID
		msgs = append(msgs, msg)

		fmt.Println(reply)
	}

	return msgs
}

func start_(user *users.Users, upd *tg.Message) tg.MessageConfig {
	reply := fmt.Sprintf("Yo @%s! I'm ipKekbot.\n"+
		"I can help you get information about the ip and store it in the database.\n\n"+
		"You can control me by sending these commands:\n\n"+
		"/checkip - get information about the ip\n"+
		"/historycheck - get history checks",
		upd.From.UserName)
	msg := tg.NewMessage(upd.Chat.ID, reply)
	msg.ReplyMarkup = users.CheckRole(user.Role)
	msg.ReplyToMessageID = upd.MessageID
	return msg
}

func checkIp_(db *gorm.DB, upd *tg.Message) tg.MessageConfig {
	reply := "Enter the ip address.\nExample: 213.180.193.1"
	msg := tg.NewMessage(upd.Chat.ID, reply)
	msg.ReplyToMessageID = upd.MessageID
	res := db.Find(&users.Users{Id: upd.From.ID}).Update("prev_msg", upd.Text)
	if res.Error != nil {
		log.Fatal(res.Error.Error())
	}
	return msg
}

func unknown_(upd *tg.Message) tg.MessageConfig {
	reply := "Unknown Command."
	msg := tg.NewMessage(upd.Chat.ID, reply)
	msg.ReplyToMessageID = upd.MessageID
	return msg
}
