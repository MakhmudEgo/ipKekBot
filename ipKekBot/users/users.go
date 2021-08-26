package users

import tg "github.com/go-telegram-bot-api/telegram-bot-api"

type Users struct {
	Id      int
	Adm     bool
	PrevMsg string
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
)
