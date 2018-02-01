package session

import (
	"time"
	"github.com/jinzhu/gorm"
)

func NewTokenRepositoryFactory() ITokenRepositoryFactory {
	return &TokenRepositoryFactory{}
}

type TokenRepositoryFactory struct{}

func (factory *TokenRepositoryFactory) New(db *gorm.DB) ITokenRepository {
	return &TokenRepository{db}
}

type TokenRepository struct {
	DB *gorm.DB
}

func (repo *TokenRepository) CreateToken(token *TokenClaims) error {
	repo.DB.Create(&token)
	
	return repo.DB.Error
}

func (repo *TokenRepository) RevokeToken(token *TokenClaims) error {
	
	repo.DB.Model(&token).Update("revoked_at", time.Now())
	
	return repo.DB.Error
}

func (repo *TokenRepository) DeleteExpiredTokens() error {

	repo.DB.Delete(TokenClaims{}, "expiry_date < ?", time.Now())
	
	return repo.DB.Error
}


func (repo *TokenRepository) FindById(id uint) (*TokenClaims, error) {
	
	var token = &TokenClaims{}	
	repo.DB.Where(&TokenClaims{ID: id}).First(&token)
		
	return token, repo.DB.Error
}

func (repo *TokenRepository) IsTokenRevoked(jti uint) bool {

	if jti == 0{
		return true
	}

	token, err := repo.FindById(jti)
	if err != nil{
		return true
	}
	
	if token.IsEmpty() == false && token.RevokedAt.IsZero() == true{
		return false;
	}
	
	return true
}