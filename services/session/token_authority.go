package session

import (
	"log"
	"fmt"
	"time"	
	"strconv"	
	"crypto/rsa"
	"go-api/domain"
	"go-api/services/user"
	"go-api/services/role"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
)

type ITokenAuthority interface {
	CreateNewSessionToken(claims domain.ITokenClaims) (string, error)
	VerifyTokenString(tokenStr string) (IToken, domain.ITokenClaims, error)
	StartSessionForUser(username string, password string) (string, error)
}

// TokenAuthority implements ITokenAuthorzity
type TokenAuthority struct {
	Options *TokenAuthorityOptions
}

type TokenAuthorityOptions struct {
	PrivateSigningKey 			*rsa.PrivateKey
	PublicSigningKey 			*rsa.PublicKey
	Database                	*gorm.DB
	TokenRepositoryFactory 		ITokenRepositoryFactory
	UserRepositoryFactory 		user.IUserRepositoryFactory
}

type TokenAuthorityClaims struct {
    UserId 	uint 		`json:"userId"`
    jwt.StandardClaims 
}

func NewTokenAuthority(options *TokenAuthorityOptions) *TokenAuthority {
	ta := TokenAuthority{options}
    
	return &ta
}

func (ta *TokenAuthority) generateNewToken(claims *TokenClaims) (*TokenAuthorityClaims, error){

	claims.ExpireAt = time.Now().Add(time.Hour * 72) // 3 days
	claims.Issuer = "todo-app"
	claims.IssuedAt = time.Now()
	
	tokenRepo := ta.TokenRepository()
	err := tokenRepo.CreateToken(claims)
	if err != nil{
		return nil, err
	}
	
	// Create the Claims
	userClaims := &TokenAuthorityClaims{
	    claims.UserID,
	    jwt.StandardClaims{
	        Id:			strconv.FormatUint(uint64(claims.ID), 10),
	        ExpiresAt: 	claims.ExpireAt.Unix(),
	        Issuer:    	claims.Issuer,
	        IssuedAt:	claims.IssuedAt.Unix(),
	    },
	}
		
	return userClaims, nil
}

func (ta *TokenAuthority) CreateNewSessionToken(claims domain.ITokenClaims) (string, error) {

	c := claims.(*TokenClaims)
	
	userClaims, err := ta.generateNewToken(c)
	if err != nil{
		return "", err
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, userClaims)
	tokenString, err := token.SignedString(ta.Options.PrivateSigningKey)

	return tokenString, err
}

func (ta *TokenAuthority) VerifyTokenString(tokenString string) (IToken, domain.ITokenClaims, error) {

	log.Println("tokenstring", tokenString)

	// Parse the token
	t, err := jwt.ParseWithClaims(tokenString, &TokenAuthorityClaims{}, func(token *jwt.Token) (interface{}, error) {
	    // since we only use the one private key to sign the tokens,
	    // we also only use its public counter part to verify
	    return ta.Options.PublicSigningKey, nil
	})
	
	if err != nil {
    	return nil, nil, err
	}
		
	var userClaims TokenClaims	
	claims, ok := t.Claims.(*TokenAuthorityClaims)
	if ok && t.Valid {
		userClaims.UserID = claims.UserId
		
		IdUint, _ := strconv.ParseUint(claims.StandardClaims.Id, 10, 64)
		
		userClaims.ID = uint(IdUint)
		userClaims.IssuedAt, _ = time.Parse(time.RFC3339, strconv.FormatInt(claims.StandardClaims.IssuedAt, 10))
		userClaims.ExpireAt, _ = time.Parse(time.RFC3339, strconv.FormatInt(claims.StandardClaims.ExpiresAt, 10))
	}else {
        log.Println("Token is not valid", err.Error())
    }
    
    log.Println("Token valid", t.Valid)
    
    tm := time.Unix(claims.StandardClaims.ExpiresAt, 0)
    
    log.Println("Expire at", tm)
    log.Println("user id", claims.UserId)

	token := NewToken(t)
	return token, &userClaims, err
}

func (ta *TokenAuthority) StartSessionForUser(username string, password string) (string, error) {
	
	if username == "" || password == "" {
		return "", fmt.Errorf("Invalid parameters")
	}
	
	userRepo := ta.UserRepository()
	registeredUser, err := userRepo.FindByUsername(&username)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	
	if registeredUser.IsEmpty() {
		var newUser = &user.User{
			Username: username,
			Roles: []role.Role{
				{ID: role.RoleUser},
			},
		}
		newUser.SetPassword(password)
		newUser.GenerateConfirmationCode()
		
		err := userRepo.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
		
		registeredUser = newUser
	}else{
	
		if registeredUser.IsCredentialsVerified(password) == false {
			return "", fmt.Errorf("Invalid username/password")
		}
	}
	
	userRepo.UpdateUserLastLogIn(registeredUser)

	tokenString, err := ta.CreateNewSessionToken(NewTokenClaims(registeredUser.GetID()))
		
	return tokenString, err
}

func (token *TokenAuthority) TokenRepository() ITokenRepository {
	return token.Options.TokenRepositoryFactory.New(token.Options.Database)
}

func (token *TokenAuthority) TokenRepositoryTransaction(tx *gorm.DB) ITokenRepository {
	return token.Options.TokenRepositoryFactory.New(tx)
}

func (token *TokenAuthority) UserRepository() user.IUserRepository {
	return token.Options.UserRepositoryFactory.New(token.Options.Database)
}

func (token *TokenAuthority) UserRepositoryTransaction(tx *gorm.DB) user.IUserRepository {
	return token.Options.UserRepositoryFactory.New(tx)
}