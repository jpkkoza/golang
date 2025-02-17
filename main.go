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

// Обновление задачи по ID (PATCH)
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи из URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Декодируем входящий JSON в структуру Task
	var updatedData Task
	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Ищем задачу в БД по ID
	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	// Обновляем только переданные поля
	if updatedData.Task != "" {
		task.Task = updatedData.Task
	}
	task.IsDone = updatedData.IsDone // Обновится, если передано в JSON

	// Сохраняем изменения
	DB.Save(&task)

	// Отправляем обновлённую задачу в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Удаление задачи по ID (DELETE)
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи из URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Проверяем, есть ли задача в базе
	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	// Удаляем задачу
	DB.Delete(&task)

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusNoContent) // 204 No Content
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
	router.HandleFunc("/api/tasks/{id}", UpdateTask).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteTask).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
