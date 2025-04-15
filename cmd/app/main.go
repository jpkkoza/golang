package main

import (
	"log"
	"pet_project_1_etap/internal/database"
	"pet_project_1_etap/internal/handlers"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/userService"
	"pet_project_1_etap/internal/web/tasks"
	"pet_project_1_etap/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Применяем миграции
	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	// ==== Task ====
	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// ==== User ====
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем task хендлеры
	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	// Регистрируем user хендлеры
	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
