package domain

import (	
	"net/http"
)

const UserInfoKey string = "mddlwr-session-user-info-key"
const TokenClaimsKey string = "mddlwr-session-token-claims-key"

func GetAuthenticatedClaimsCtx(r *http.Request) ITokenClaims {
	if claim := r.Context().Value(TokenClaimsKey); claim != nil {
		return claim.(ITokenClaims)
	}
	return nil
}

func GetAuthenticatedUserCtx(r *http.Request) IUser {
	if user := r.Context().Value(UserInfoKey); user != nil {
		return user.(IUser)
	}
	return nil
}

