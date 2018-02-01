package session

import (
	"github.com/jinzhu/gorm"
)

type ITokenRepositoryFactory interface {
	New(db *gorm.DB) ITokenRepository
}
type ITokenRepository interface {
	CreateToken(token *TokenClaims) error
	RevokeToken(token *TokenClaims) error
	DeleteExpiredTokens() error
	IsTokenRevoked(jti uint) bool
	FindById(id uint) (*TokenClaims, error)
}