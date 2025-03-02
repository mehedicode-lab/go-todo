package repositories

import "go-todo/internal/domain/models"

type TodoRepository interface {
	Create(todo *models.Todo) error
	GetAll() ([]models.Todo, error)
	GetByID(id uint) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id uint) error
}
