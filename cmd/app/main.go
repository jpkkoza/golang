package main

import (
	"log"
	"net/http"
	"pet_project_1_etap/internal/database"
	"pet_project_1_etap/internal/handlers"
	"pet_project_1_etap/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализируем базу данных
	database.InitDB()

	// Создаем репозиторий, сервис и хендлер
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Настройка роутера
	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	// Запуск сервера с обработкой ошибки
	log.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
