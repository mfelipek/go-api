package user

import (
	"github.com/jinzhu/gorm"
)

type IUserRepositoryFactory interface {
	New(db *gorm.DB) IUserRepository
}

type IUserRepository interface {
	CreateUser(user *User) (err error)
	FindByUsername(username *string) (*User, error)
	FindById(id uint) (*User, error)
	UpdateUserLastLogIn(user *User) (err error)	
}