package domain

type IUser interface {
	GetID() uint
	IsValid() bool
	IsCodeVerified(code string) bool
	IsCredentialsVerified(password string) bool
	SetPassword(password string) error
	GenerateConfirmationCode()
	HasRole(r IRole) bool
	IsEmpty() bool
}

type IRole interface{}

