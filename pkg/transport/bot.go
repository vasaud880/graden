package transport

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vasaud880/graden/pkg/service"
)

type Bot struct {
	botAPI      *tgbotapi.BotAPI
	planService *service.PlanService
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		ctx := context.Background()

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				b.sendMessage(chatID, "Добро пожаловать в Градостроитель! Используйте команды для управления.")
			case "newplan":
				_, err := b.planService.CreatePlan(ctx, chatID)
				if err != nil {
					b.sendMessage(chatID, "Ошибка при создании плана.")
				} else {
					b.sendMessage(chatID, "Создан новый градостроительный план. Введите название города.")
				}
			case "viewplan":
				plan, err := b.planService.GetPlan(ctx, chatID)
				if err != nil {
					b.sendMessage(chatID, "У вас нет активного плана. Используйте /newplan для создания нового.")
				} else {
					b.sendMessage(chatID, fmt.Sprintf("Название: %s\nОписание: %s\nПлощадь: %.2f км²\nНаселение: %d", plan.Name, plan.Description, plan.Area, plan.Population))
				}
			case "deleteplan":
				err := b.planService.DeletePlan(ctx, chatID)
				if err != nil {
					b.sendMessage(chatID, "Ошибка при удалении плана.")
				} else {
					b.sendMessage(chatID, "План успешно удален.")
				}
			case "help":
				b.sendMessage(chatID, "Доступные команды:\n/newplan - создать новый план\n/viewplan - просмотреть текущий план\n/deleteplan - удалить план")
			default:
				b.sendMessage(chatID, "Неизвестная команда. Используйте /help для списка команд.")
			}
		} else {
			// Обработка текстовых сообщений для заполнения плана
			plan, err := b.planService.GetPlan(ctx, chatID)
			if err != nil {
				b.sendMessage(chatID, "Используйте /newplan для создания нового плана.")
				continue
			}

			if plan.Name == "" {
				b.planService.UpdatePlanName(ctx, chatID, update.Message.Text)
				b.sendMessage(chatID, "Название города установлено. Введите описание.")
			} else if plan.Description == "" {
				b.planService.UpdatePlanDescription(ctx, chatID, update.Message.Text)
				b.sendMessage(chatID, "Описание установлено. Введите площадь города в км².")
			} else if plan.Area == 0 {
				area, err := strconv.ParseFloat(update.Message.Text, 64)
				if err != nil {
					b.sendMessage(chatID, "Пожалуйста, введите корректное число для площади.")
				} else {
					b.planService.UpdatePlanArea(ctx, chatID, area)
					b.sendMessage(chatID, "Площадь установлена. Введите население города.")
				}
			} else if plan.Population == 0 {
				population, err := strconv.Atoi(update.Message.Text)
				if err != nil {
					b.sendMessage(chatID, "Пожалуйста, введите корректное число для населения.")
				} else {
					b.planService.UpdatePlanPopulation(ctx, chatID, population)
					b.sendMessage(chatID, "Население установлено. Используйте /viewplan для просмотра вашего плана.")
				}
			}
		}
	}
}

func (b *Bot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.botAPI.Send(msg)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func NewBot(token string, planService *service.PlanService) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{botAPI: botAPI, planService: planService}, nil
}
