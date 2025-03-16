package main

import (
	"log"

	"github.com/vasaud880/graden/pkg/repository"
	"github.com/vasaud880/graden/pkg/service"
	"github.com/vasaud880/graden/pkg/transport"
)

func main() {
	repo := repository.NewPlanRepository()
	planService := service.NewPlanService(repo)

	bot, err := transport.NewBot("YOUR_TELEGRAM_BOT_API_TOKEN", planService)
	if err != nil {
		log.Panic(err)
	}

	bot.Start()
}
