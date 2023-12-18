package handlers

import (
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotHandlers struct {
	StartHandler        func(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages)
	SendMenuHandler     func(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages)
	ReadLogsHandler     func(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages)
	DescGroupsHandler   func(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages)
	GetIdHandler        func(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages)
	CalendarWeekHandler func(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages)
}
