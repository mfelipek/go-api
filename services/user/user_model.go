package user

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/twinj/uuid"
	"go-api/domain"
	"go-api/services/role"
)

type User struct {
	ID        		uint 			`gorm:"primary_key"`
	Username		string 			`json:"username" gorm:"size:30"`
	LastLogin		time.Time 		`json:"lastLogin,omitempty"`
	
	Roles         []role.Role		`gorm:"many2many:user_role;"`
	
	// fields are not exported to JSON
	ConfirmationCode string `json:"-" gorm:"size:120"`
	HashedPassword   string `json:"-" gorm:"size:200"`
}

func (User) TableName() string {
  return "user"
}

func (user *User) GetID() uint {
	return user.ID
}

func (user *User) IsValid() bool {
	return len(user.Username) > 0
}

func (user *User) IsCodeVerified(code string) bool {
	return (user.ConfirmationCode == code)
}

func (user *User) IsCredentialsVerified(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return (err == nil)
}

func (user *User) SetPassword(password string) error {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashedPassword = string(hashedPassword)
	return nil
}

func (user *User) GenerateConfirmationCode() {
	user.ConfirmationCode = generateNewUniqueCode()
}

func generateNewUniqueCode() string {
	// set code format
	uuid.SwitchFormat(uuid.FormatCanonical)
	return uuid.NewV4().String()
}

func (user *User) IsEmpty() bool {
	return user.ID == 0
}

func (user *User) HasRole(r domain.IRole) bool {
	role := r.(role.Role)
	for _, a := range user.Roles {
		if a.ID == role.ID {
			return true
		}
	}
	return false
}