package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"

	"github.com/vasaud880/graden/pkg/repository"
	"github.com/vasaud880/graden/pkg/service"
	"github.com/vasaud880/graden/pkg/transport"
)

func main() {
	// Подключение к PostgreSQL
	dbURL := "postgres://vasaud880:12345@localhost:5432/graden"
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Инициализация репозитория и сервиса
	repo := repository.NewPlanRepository(db)
	planService := service.NewPlanService(repo)

	// Запуск бота
	bot, err := transport.NewBot(os.Getenv("7769665201:AAGRKKII2_uuAeYlrl0eUL8kalJ3ULvSHMI"), planService)
	if err != nil {
		log.Panic(err)
	}

	bot.Start()
}
