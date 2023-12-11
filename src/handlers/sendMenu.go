package handlers

import (
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"
	"slices"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMenu(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages) {
	userChatID := update.Message.Chat.ID
	var (
		commandsText string
		sb           strings.Builder
	)

	switch {
	case fmt.Sprintf("%d", userChatID) == config.AdminChatIDEvg:
		for key, value := range utils.AdminCommands {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
		commandsText = sb.String()
	case fmt.Sprintf("%d", userChatID) == config.AdminChatIDNat:
		for key, value := range utils.AdminCommands {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
		commandsText = sb.String()
	case slices.Contains(config.CoachChatIDs, fmt.Sprintf("%d", userChatID)):
		for key, value := range utils.CoachCommands {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
		commandsText = sb.String()
	default:
		for key, value := range utils.UserCommands {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
		commandsText = sb.String()
	}

	msg := tgbotapi.NewMessage(userChatID, commandsText)
	bot.Send(msg)
}
