package domain

import (
	"context"
	"net/http"
)

type ACLHandlerFunc func(*http.Request, IUser) (bool, string)

type ACLMap map[string]ACLHandlerFunc

func (m *ACLMap) Append(maps ...*ACLMap) ACLMap {
	res := ACLMap{}
	// copy current map
	for k, v := range *m {
		res[k] = v
	}
	for _, _maps := range maps {
		for k, v := range *_maps {
			res[k] = v
		}
	}
	return res
}

type IAccessController interface {
	Add(*ACLMap)
	AddHandler(name string, handler ACLHandlerFunc)
	HasAction(string) bool
	IsHTTPRequestAuthorized(req *http.Request, ctx context.Context, action string, user IUser) (bool, string)
	NewContextHandler(string, http.HandlerFunc) http.HandlerFunc
}