package site

import (
	"go-api/domain"
	"net/http"
)

func (resource *SiteResource) HandleIndexACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous access
	return true, ""
}

func (resource *SiteResource) HandleErrorACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous access
	return true, ""
}