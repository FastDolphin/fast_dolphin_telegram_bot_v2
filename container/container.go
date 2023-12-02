package container

import (
	"fast_dolphin_telegram_bot/src/handlers"
	"fast_dolphin_telegram_bot/src/utils"
	"fmt"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	di := dig.New()
	loadConfig(di)

	must(di.Provide(func(cfg *utils.Config) *handlers.BotHandlers {
		return &handlers.BotHandlers{
			StartHandler: handlers.Start, // Make sure this matches the new signature
			// Initialize other handlers here
		}
	}))

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
