package handlers

import (
	"fast_dolphin_telegram_bot/logger"
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"
	"strings"
	"time"

	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CalendarWeek(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *utils.Config, messages *messageUtils.Messages) {
	userChatID := update.Message.Chat.ID
	chatIdStr := fmt.Sprintf("%d", userChatID)

	if utils.IsAdmin(chatIdStr) || utils.IsCoach(chatIdStr) {
		_, calWeek := time.Now().ISOWeek()
		calWeekStr := strconv.Itoa(calWeek)
		text := strings.Replace(messages.CalendarWeek, "{calender_week}", calWeekStr, 1)
		msg := tgbotapi.NewMessage(userChatID, text)
		bot.Send(msg)
	} else {
		logger.Log.Warn(fmt.Sprintf("Unauthorized log access attempt by user %d.", userChatID))
		msg := tgbotapi.NewMessage(userChatID, messages.NotAllowed)
		bot.Send(msg)
	}
}
