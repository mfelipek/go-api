package todo

import (
	"time"
	"go-api/services/user"
)

const (
	TodoStatusPending   = "pending"
	TodoStatusCompleted   = "completed"
	TodoStatusInactive  = "inactive"
)

type Todo struct {
	ID        	uint 		`json:"id" gorm:"primary_key"`
	UserID   	uint
	User        user.User	`json:"-"`
    Name		string 		`json:"name"`
    Status		string 		`json:"status,omitempty"`
    Due			time.Time 	`json:"due,omitempty"`
}

func (Todo) TableName() string {
  return "todo"
}

type ITodo interface {
	GetID() uint
	IsValid() bool
}

type NewTodo struct {
	Name 		string	`json:"name,omitempty"`
}

type TodoList []*Todo

func (todo *Todo) GetID() uint {
	return todo.ID
}

func (todo *Todo) IsValid() bool {
	
	return len(todo.Name) > 0
}