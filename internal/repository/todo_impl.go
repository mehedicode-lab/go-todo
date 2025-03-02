package repository

import (
	"go-todo/internal/domain/models"
	"go-todo/internal/domain/repositories"
	"gorm.io/gorm"
)

type TodoRepositoryImp struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repositories.TodoRepository {
	return &TodoRepositoryImp{
		db: db,
	}
}

func (r *TodoRepositoryImp) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepositoryImp) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepositoryImp) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepositoryImp) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepositoryImp) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
