package session

import (
	"go-api/domain"
	"net/http"
)

func (resource *SessionResource) HandleGetSessionACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		return false, ""
	}
	return true, ""
}

func (resource *SessionResource) HandleCreateSessionACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous access
	return true, ""
}

func (resource *SessionResource) HandleDeleteSessionACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous access
	return true, ""
}