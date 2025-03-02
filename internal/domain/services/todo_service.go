package services

import (
	"go-todo/internal/domain/models"
	"go-todo/internal/domain/repositories"
)

type TodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) Create(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) GetByID(id uint) (*models.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) Update(todo *models.Todo) error {
	return s.repo.Update(todo)
}

func (s *TodoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
