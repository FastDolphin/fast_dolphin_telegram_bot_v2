package container

import (
	"fast_dolphin_telegram_bot/src/handlers"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	di := dig.New()
	loadConfig(di)
	loadBotAPI(di)
	loadHandlers(di)
	return di
}

func MustInvoke(container *dig.Container, function interface{}, opts ...dig.InvokeOption) {
	must(container.Invoke(function, opts...))
}

// must is a helper function for error handling in DI container setup
func must(err error) {
	if err != nil {
		panic(fmt.Sprintf("failed to initialize DI: %s", err))
	}
}

func loadConfig(container *dig.Container) {
	must(container.Provide(utils.New))
}

func loadBotAPI(container *dig.Container) {
	must(container.Provide(func(cfg *utils.Config) (*tgbotapi.BotAPI, error) {
		bot, err := tgbotapi.NewBotAPI(cfg.Token)
		if err != nil {
			return nil, err
		}
		return bot, nil
	}))
}

func loadHandlers(container *dig.Container) {
	must(container.Provide(func(cfg *utils.Config) *handlers.BotHandlers {
		return &handlers.BotHandlers{
			StartHandler:    handlers.Start,
			SendMenuHandler: handlers.SendMenu,
			// Initialize other handlers here
		}
	}))
}
