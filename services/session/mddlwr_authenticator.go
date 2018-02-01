package session

import (
	"log"
	"context"
	"go-api/domain"
	"net/http"
	"strings"
)

func NewAuthenticator(resource *SessionResource) *Authenticator {
	return &Authenticator{resource}
}

// SessionsAuthenticator implements IMiddleware
type Authenticator struct {
	resource *SessionResource
}

// Handler authenticates a session token in the Authorization header
func (auth *Authenticator) Handler(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	resource := auth.resource
	ctx := resource.Context()

	authHeaderString := req.Header.Get("Authorization")
	if authHeaderString != "" {
		tokens := strings.Split(authHeaderString, " ")
		if len(tokens) != 2 || (len(tokens) > 0 && strings.ToUpper(tokens[0]) != "BEARER") {
			resource.RenderUnauthorizedError(w, req, "Invalid format, expected Authorization: Bearer {token}")
			return
		}
		tokenString := strings.TrimSpace(tokens[1])
		t, c, err := resource.TokenAuthority.VerifyTokenString(tokenString)
		if err != nil {
			resource.RenderUnauthorizedError(w, req, "Unable to verify token string" + err.Error())
			return
		}
		token := t.(*Token)
		claims := c.(*TokenClaims)
		if !token.Valid {
			resource.RenderUnauthorizedError(w, req, "Token is invalid")
			return
		}

		// Check that the token was not previously revoked
		// TODO: Possible optimization, use Redis
		tokenRepo := resource.TokenRepository()
		
		if tokenRepo.IsTokenRevoked(claims.ID) {
			resource.RenderUnauthorizedError(w, req, "Token has been revoked")
			return
		}

		// retrieve user object and store it in current session request context
		// this `user` object will be used by the AccessController middleware
		userRepo := resource.UserRepository()
		
		user, err := userRepo.FindById(claims.UserID)
		if err != nil {
			log.Println("userRepo.FindById ", err.Error())
		
			// `user` = nil indicates that current authentication failed
			user = nil
		}

		// add claims and current user object to session request context
		ctx = setAuthenticatedClaimsCtx(req, claims)
		ctx = SetCurrentUserCtx(req, user)
	}

	next(w, req.WithContext(ctx))
}

func setAuthenticatedClaimsCtx(r *http.Request, claim domain.ITokenClaims) (context.Context) {
	return context.WithValue(r.Context(), domain.TokenClaimsKey, claim)
}

func SetCurrentUserCtx(r *http.Request, user domain.IUser) (context.Context) {
	return context.WithValue(r.Context(), domain.UserInfoKey, user)
}