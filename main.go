package main

import (
	"fast_dolphin_telegram_bot/container"
	messageUtils "fast_dolphin_telegram_bot/messages"
	"fast_dolphin_telegram_bot/src/handlers"
	commonUtils "fast_dolphin_telegram_bot/src/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	di := container.CreateContainer()

	container.MustInvoke(di, func(
		bot *tgbotapi.BotAPI,
		botHandlers *handlers.BotHandlers,
		config *commonUtils.Config,
	) {
		messagesStruct, err := messageUtils.LoadMessages(config.MessagesPath)
		if err != nil {
			panic(err)
		}

		// Setup update configuration and channel
		updateConfig := tgbotapi.NewUpdate(0)
		updateConfig.Timeout = 60
		updates, _ := bot.GetUpdatesChan(updateConfig)

		// Listen for updates and handle commands
		for update := range updates {
			if update.Message == nil {
				continue
			}

			switch update.Message.Command() {
			case "start":
				botHandlers.StartHandler(bot, update, config, messagesStruct)
			case "menu":
				botHandlers.SendMenuHandler(bot, update, config, messagesStruct)
			case "read_logs":
				botHandlers.ReadLogsHandler(bot, update, config, messagesStruct)
			case "get_id":
				botHandlers.GetIdHandler(bot, update, config, messagesStruct)
			case "desc_groups":
				botHandlers.DescGroupsHandler(bot, update, config, messagesStruct)
				// Add other command cases as needed
				// default:
				// 	botHandlers.StartHandler(bot, update, config, messagesStruct)
			}
		}
	})
}
