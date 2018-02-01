package user

import (
	"time"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepositoryFactory() IUserRepositoryFactory {
	return &UserRepositoryFactory{}
}

type UserRepositoryFactory struct{}

func (factory *UserRepositoryFactory) New(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *User) (err error) {
	
	repo.DB.Create(&user)
	
	return repo.DB.Error
}

func (repo *UserRepository) FindByUsername(username *string) (*User, error) {

	var user = &User{}
	repo.DB.Where(&User{Username: *username}).First(&user)
		
	return user, repo.DB.Error
}

func (repo *UserRepository) FindById(id uint) (*User, error) {

	var userDb = &User{}
	repo.DB.Where(&User{ID: id}).First(&userDb)
		
	return userDb, repo.DB.Error
}

func (repo *UserRepository) UpdateUserLastLogIn(user *User) (err error) {
	
	repo.DB.Model(&user).Update("last_login_date", time.Now())
	
	return repo.DB.Error
}