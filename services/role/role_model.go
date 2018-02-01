package role

import (
)

// Types of Roles
const (
	RoleAdmin uint = 1
	RoleUser  uint = 2
)

type Role struct {
	ID        		uint 			`gorm:"primary_key"`
	Description		string 			`json:"description" gorm:"size:30"`
}

func (Role) TableName() string {
  return "role"
}