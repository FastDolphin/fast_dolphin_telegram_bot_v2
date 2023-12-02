package handlers

import (
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages) {
	userChatID := update.Message.Chat.ID
	var greetingMessage string

	switch {
	case fmt.Sprintf("%d", userChatID) == config.AdminChatIDEvg:
		greetingMessage = fmt.Sprintf(messages.StartMessageAdmin, config.AdminNameEvg)
	case fmt.Sprintf("%d", userChatID) == config.AdminChatIDNat:
		greetingMessage = fmt.Sprintf(messages.StartMessageAdmin, config.AdminNameNat)
	case contains(config.CoachChatIDs, fmt.Sprintf("%d", userChatID)):
		greetingMessage = messages.StartMessageCoach
	default:
		greetingMessage = messages.StartMessage
	}

	msg := tgbotapi.NewMessage(userChatID, greetingMessage)
	bot.Send(msg)
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
