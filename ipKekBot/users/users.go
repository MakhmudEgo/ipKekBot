package users

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	UserR = iota
	AdminR
	CreatorR
)

type UserHistory struct {
	IpsId  uint `gorm:"column:ips_id"`
	Time   time.Time
	UserId int `gorm:"column:user_id"`
}

type Users struct {
	Id       int
	Username string
	Role     int
	PrevMsg  string
}

var UserBtn = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/checkip"),
		tg.NewKeyboardButton("/historycheck"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("test"),
		tg.NewKeyboardButton("213.180.193.1"),
		tg.NewKeyboardButton("195.133.239.83"),
		tg.NewKeyboardButton("209.85.229.104"),
	),
)

var AdminBtn = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/checkip"),
		tg.NewKeyboardButton("/historycheck"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/sendmessage"),
		tg.NewKeyboardButton("/addnewadmin"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/getlistchecksipuser"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("test"),
		tg.NewKeyboardButton("213.180.193.1"),
		tg.NewKeyboardButton("195.133.239.83"),
		tg.NewKeyboardButton("209.85.229.104"),
	),
)

var CreatorBtn = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/checkip"),
		tg.NewKeyboardButton("/historycheck"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/sendmessage"),
		tg.NewKeyboardButton("/addnewadmin"),
		tg.NewKeyboardButton("/deleteadmin"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/getlistchecksipuser"),
	),

	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("test"),
		tg.NewKeyboardButton("213.180.193.1"),
		tg.NewKeyboardButton("195.133.239.83"),
		tg.NewKeyboardButton("209.85.229.104"),
	),
)

func CheckRole(role int) interface{} {
	switch role {
	case CreatorR:
		return CreatorBtn
	case AdminR:
		return AdminBtn
	case UserR:
		return UserBtn
	}
	log.Fatal("unknown role")
	return nil
}
