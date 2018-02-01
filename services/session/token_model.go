package session

import (
	"time"
	"go-api/services/user"
	"github.com/dgrijalva/jwt-go"
)



type IToken interface {
	IsValid() bool
}

func NewTokenClaims(userID uint) *TokenClaims {
	return &TokenClaims{UserID: userID}
}

type TokenClaims struct {
	ID      		uint			`gorm:"primary_key"`
	UserID   		uint
	User        	user.User
	Issuer			string
	ExpireAt 		time.Time
	IssuedAt 		time.Time
	RevokedAt	 	time.Time
}

func (TokenClaims) TableName() string {
  return "access_token"
}

func (claim *TokenClaims) GetJTI() uint {
	return claim.ID
}

func (claim *TokenClaims) GetExpireAt() time.Time {
	return claim.ExpireAt
}

func (token *TokenClaims) IsEmpty() bool {
	return token.ID == 0
}

func NewToken(token *jwt.Token) *Token {
	return &Token{token}
}

type Token struct {
	*jwt.Token
}

func (token *Token) IsValid() bool {
	return token.Valid
}