package todo

import (
	"github.com/jinzhu/gorm"
)

type ITodoRepositoryFactory interface {
	New(db *gorm.DB) ITodoRepository
}

type ITodoRepository interface {
	CreateTodo(todo ITodo) (err error)
	FindAll(userId uint) ([]*Todo, error)
}