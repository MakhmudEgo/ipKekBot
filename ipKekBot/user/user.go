package user

import tg "github.com/go-telegram-bot-api/telegram-bot-api"

var userBtn = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
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
