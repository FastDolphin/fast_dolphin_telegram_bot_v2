package handlers

import (
	"fast_dolphin_telegram_bot/logger"
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DescGroups(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages) {
	userChatID := update.Message.Chat.ID
	chatIdStr := fmt.Sprintf("%d", userChatID)

	if utils.IsAdmin(chatIdStr) || utils.IsCoach(chatIdStr) {
		messageWithTitle := messages.DescGroupMesage
		maxMessageLength := 4096 // Assuming MAX_MESSAGE_LENGTH is 4096

		dataChunks := make([]string, 0)
		for i := 0; i < len(messageWithTitle); i += maxMessageLength {
			end := i + maxMessageLength
			if end > len(messageWithTitle) {
				end = len(messageWithTitle)
			}
			dataChunks = append(dataChunks, messageWithTitle[i:end])
		}

		for _, chunk := range dataChunks {
			msg := tgbotapi.NewMessage(userChatID, chunk)
			bot.Send(msg)
		}
	} else {
		logger.Log.Warn(fmt.Sprintf("Unauthorized log access attempt by user %d.", userChatID))
		msg := tgbotapi.NewMessage(userChatID, messages.NotAllowed)
		bot.Send(msg)
	}
}
