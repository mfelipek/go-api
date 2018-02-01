package todo

import (
	"go-api/domain"
	"net/http"
)

func (resource *TodoResource) HandleTodoListACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		return false, ""
	}
	return true, ""
}

func (resource *TodoResource) HandleTodoCreateACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		return false, ""
	}
	return true, ""
}