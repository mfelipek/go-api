package todo

import (
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepositoryFactory() ITodoRepositoryFactory {
	return &TodoRepositoryFactory{}
}

type TodoRepositoryFactory struct{}

func (factory *TodoRepositoryFactory) New(db *gorm.DB) ITodoRepository {
	return &TodoRepository{db}
}

func (repo *TodoRepository) FindAll(userId uint) ([]*Todo, error) {

	var todoList []*Todo
	
	repo.DB.Where(&Todo{UserID: userId}).Find(&todoList)
	
	return todoList, repo.DB.Error
}

func (repo *TodoRepository) CreateTodo(_todo ITodo) (err error) {

	todo := _todo.(*Todo)
	todo.Status = TodoStatusPending
		
	repo.DB.Create(&todo)
	
	return repo.DB.Error
}