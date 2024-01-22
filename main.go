package main

import (
	bardService "face-id-detection/internal/modules/bard/services"
	telegramService "face-id-detection/internal/modules/telegram/service"
	"face-id-detection/pkg/logger"
)

func main() {
	logger.Start()
	logger.Info("Started...")
	bardService.Init()

	// if err := env.Start(); err != nil {
	// 	panic(err.Error())
	// }

	telegramService.Send()
}
