package handlers

import (
	"fast_dolphin_telegram_bot/logger"
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetId(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages) {
	userChatID := update.Message.Chat.ID
	userName := update.Message.From.UserName

	logger.Log.Info("Triggering get_id for user " + userName + "...")

	message := fmt.Sprintf(messages.GetId, userName, userChatID)

	msg := tgbotapi.NewMessage(userChatID, message)

	logger.Log.Info("Triggered get_id for user " + userName + " and got ID " + fmt.Sprintf("%d", userChatID))

	bot.Send(msg)
}
