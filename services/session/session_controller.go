package session

import (
	"encoding/json"
	"fmt"
	"go-api/domain"
	"log"
	"net/http"
)

type GetSessionResponse_v0 struct {
	User    domain.IUser `json:"user"`
	Success bool         `json:"success"`
	Message string       `json:"message"`
}
type CreateSessionRequest_v0 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CreateSessionResponse_v0 struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type DeleteSessionResponse_v0 struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type ErrorResponse_v0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

func (resource *SessionResource) DecodeRequestBody(w http.ResponseWriter, req *http.Request, target interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		resource.RenderError(w, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return err
	}
	return nil
}

func (resource *SessionResource) RenderError(w http.ResponseWriter, req *http.Request, status int, message string) {
	resource.Renderer.JSON(w, status, ErrorResponse_v0{
		Message: message,
		Success: false,
	})
}

func (resource *SessionResource) RenderUnauthorizedError(w http.ResponseWriter, req *http.Request, message string) {
	resource.Renderer.JSON(w, http.StatusUnauthorized, ErrorResponse_v0{
		Message: message,
		Success: false,
	})
}

// HandleGetSession_v0 Get session details
func (resource *SessionResource) HandleGetSession_v0(w http.ResponseWriter, req *http.Request) {
	user := domain.GetAuthenticatedUserCtx(req)

	resource.Renderer.JSON(w, http.StatusOK, GetSessionResponse_v0{
		User:    user,
		Success: true,
		Message: "Session details retrieved",
	})
}

// HandleCreateSession_v0 verify user's credentials and generates a JWT token if valid
func (resource *SessionResource) HandleCreateSession_v0(w http.ResponseWriter, req *http.Request) {
	ta := resource.TokenAuthority

	var body CreateSessionRequest_v0
	err := resource.DecodeRequestBody(w, req, &body)
	if err != nil {
		return
	}	
	
	tokenString, err := ta.StartSessionForUser(body.Username, body.Password)
	if err != nil {
		resource.RenderError(w, req, http.StatusBadRequest, "Error creating session token" + err.Error())
		return
	}
	
	resource.Renderer.JSON(w, http.StatusCreated, CreateSessionResponse_v0{
		Token:   tokenString,
		Success: true,
		Message: "Session token created",
	})
}

// HandleDeleteSession_v0 invalidates a session token
func (resource *SessionResource) HandleDeleteSession_v0(w http.ResponseWriter, req *http.Request) {
	
	claims := domain.GetAuthenticatedClaimsCtx(req)
	//	hooks := ctx.GetControllerHooksMapCtx(req)

	if claims == nil || claims.GetJTI() == 0 {
		// simply return because we can't blacklist a token without identifier
		resource.Renderer.JSON(w, http.StatusOK, DeleteSessionResponse_v0{
			Success: true,
			Message: "Session removed",
		})
		return
	}
	
	repo := resource.TokenRepository()
	err := repo.RevokeToken(&TokenClaims{
			ID:	claims.GetJTI(),
	})
	
	if err != nil {
		log.Println("HandleDeleteSession_v0: Failed to revoke token", err.Error())
	}

	resource.Renderer.JSON(w, http.StatusOK, DeleteSessionResponse_v0{
		Success: true,
		Message: "Session removed",
	})
}