package domain

import (
	"time"
)

type ITokenClaims interface {
	GetJTI() uint
	GetExpireAt() time.Time
	IsEmpty() bool
}

type IToken interface {
	IsValid() bool
}