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
	Query  string
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
		tg.NewKeyboardButton("/check_ip"),
		tg.NewKeyboardButton("/history_check"),
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
		tg.NewKeyboardButton("/check_ip"),
		tg.NewKeyboardButton("/history_check"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/send_message"),
		tg.NewKeyboardButton("/add_new_admin"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/get_history_by_tg"),
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
		tg.NewKeyboardButton("/check_ip"),
		tg.NewKeyboardButton("/history_check"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/send_message"),
		tg.NewKeyboardButton("/add_new_admin"),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton("/delete_admin"),
		tg.NewKeyboardButton("/get_history_by_tg"),
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
