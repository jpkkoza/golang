package handlers

import (
	"context"

	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

// func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	// Получаем ID из URL
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Некорректный ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Декодируем тело запроса в структуру Task
// 	var updatedTask taskService.Task
// 	err = json.NewDecoder(r.Body).Decode(&updatedTask)
// 	if err != nil {
// 		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// Вызываем сервисный метод для обновления задачи
// 	task, err := h.Service.UpdateTaskByID(uint(id), updatedTask)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Отправляем обновленную задачу клиенту
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(task)
// }

// func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	// Получаем ID из URL
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Некорректный ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Вызываем сервисный метод для удаления задачи
// 	err = h.Service.DeleteTaskByID(uint(id))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

//		// Отправляем пустой успешный ответ
//		w.WriteHeader(http.StatusNoContent)
//	}
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id

	updated := request.Body

	task := taskService.Task{
		Task:   *updated.Task,
		IsDone: *updated.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), task)
	if err != nil {
		return nil, err
	}

	response := tasks.Task{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return tasks.PatchTasksId200JSONResponse(response), nil

}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil // 204 No Content
}
