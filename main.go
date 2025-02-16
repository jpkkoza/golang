package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Создание задачи (POST)
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Запись задачи в БД
	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, "Ошибка записи в БД", http.StatusInternalServerError)
		return
	}

	// Отправка ответа клиенту
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// Получение всех tasks (GET)
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []string // Слайс для хранения всех значений task

	// Выбираем только поле Task из всех записей
	var allTasks []Task
	DB.Select("task").Find(&allTasks)

	// Заполняем слайс строками task
	for _, t := range allTasks {
		tasks = append(tasks, t.Task)
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	// Инициализация БД
	InitDB()

	// Автоматическая миграция модели Task
	DB.AutoMigrate(&Task{})

	// Настройка роутера
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")

	// Запуск сервера
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
