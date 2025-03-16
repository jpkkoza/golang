package taskService

import (
	"fmt"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updatedTask Task) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}

	err := r.db.Model(&task).Updates(updatedTask).Error
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return nil
}
