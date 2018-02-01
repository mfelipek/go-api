package server

import (
	"log"
	"context"
	"go-api/domain"
	"net/http"
	"gopkg.in/unrolled/render.v1"
)

const defaultForbiddenAccessMessage = "Forbidden (403)"
const defaultOKAccessMessage = "OK"

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

// TODO: Currently, AccessController only acts as a gateway for endpoints on router level. Build AC to handler other aspects of ACL
func NewAccessController(ctx context.Context, renderer *render.Render) *AccessController {
	return &AccessController{domain.ACLMap{}, ctx, renderer}
}

// implements IAccessController
type AccessController struct {
	ACLMap   domain.ACLMap
	ctx      context.Context
	renderer *render.Render
}

func (ac *AccessController) Add(aclMap *domain.ACLMap) {
	ac.ACLMap = ac.ACLMap.Append(aclMap)
}

func (ac *AccessController) AddHandler(action string, handler domain.ACLHandlerFunc) {
	ac.ACLMap[action] = handler
}

func (ac *AccessController) HasAction(action string) bool {
	fn := ac.ACLMap[action]
	return (fn != nil)
}

func (ac *AccessController) IsHTTPRequestAuthorized(req *http.Request, ctx context.Context, action string, user domain.IUser) (bool, string) {
	fn := ac.ACLMap[action]
	if fn == nil {
		// by default, if acl action/handler is not defined, request is not authorized
		return false, defaultForbiddenAccessMessage
	}

	result, message := fn(req, user)
	if result && message == "" {
		message = defaultOKAccessMessage
	}
	if !result && message == "" {
		message = defaultForbiddenAccessMessage
	}
	return result, message
}

func (ac *AccessController) NewContextHandler(action string, next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
	
		user := domain.GetAuthenticatedUserCtx(req)
		
		log.Println("user ctx ", user)
		
		// `user` might be `nil` if has not authenticated.
		// ACL might want to allow anonymous / non-authenticated access (for login, e.g)
		result, message := ac.IsHTTPRequestAuthorized(req, ac.ctx, action, user)
		if !result {
			ac.renderer.JSON(w, http.StatusForbidden, ErrorResponse{
				Message: message,
				Success: false,
			})
			return
		}

		next(w, req)
	}
}