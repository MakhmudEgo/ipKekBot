package users

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

type Command interface {
	handleCheckIp(db *gorm.DB, message *tg.Message, query *string) tg.MessageConfig
}

type Users struct {
	Id      int
	Adm     bool
	PrevMsg string
}

type User struct {
	user Users
}

type Admin struct {
	admin Users
}

type Creator struct {
	creator Users
}

func Execute(c Command) {
	switch c.(type) {
	case *User:
		//u, _ := c.(*User)
		//u.handleCheckIp()
	}
}

func (u *User) handleCheckIp(db *gorm.DB, message *tg.Message, query *string) tg.MessageConfig {
	return tg.MessageConfig{}
}

func (u *Admin) handleCheckIp(db *gorm.DB, message *tg.Message, query *string) tg.MessageConfig {
	return tg.MessageConfig{}
}

var UserBtn = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/checkip"),
		tg.NewKeyboardButton("/historycheck"),
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
