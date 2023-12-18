package handlers

import (
	"fast_dolphin_telegram_bot/logger"
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ReadLogs(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages) {
	userChatID := update.Message.Chat.ID
	chatIdStr := fmt.Sprintf("%d", userChatID)

	if chatIdStr == config.AdminChatIDEvg || chatIdStr == config.AdminChatIDNat {
		logData, err := os.ReadFile(logger.FULL_LOG_PATH)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("Failed to open log file: %s.", err))
			return
		}

		// If the data is large, send it as a file, otherwise, send it directly
		if len(logData) > 3000 {
			file, err := os.Open(logger.FULL_LOG_PATH)
			if err != nil {
				logger.Log.Error(fmt.Sprintf("Failed to open log file: %s.", err))
				return
			}

			msg := tgbotapi.NewDocumentUpload(userChatID, tgbotapi.FileReader{
				Name:   logger.OUTPUT_LOG_FILE_NAME,
				Reader: file,
				Size:   -1, // Set to -1 to let Telegram determine the size
			})
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(userChatID, string(logData))
			bot.Send(msg)
		}
		logger.Log.Info(fmt.Sprintf("Log data accessed by admin %d.", update.Message.From.ID))
	} else {
		logger.Log.Warn(fmt.Sprintf("Unauthorized log access attempt by user %d.", update.Message.From.ID))
		msg := tgbotapi.NewMessage(userChatID, messages.NotAllowed)
		bot.Send(msg)
	}
}
