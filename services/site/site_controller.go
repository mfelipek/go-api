package site

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest_v0 struct {
	Username string `json:"username"`
}

func (resource *SiteResource) DecodeRequestBody(w http.ResponseWriter, req *http.Request, target interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		resource.RenderError(w, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return err
	}
	return nil
}

func (resource *SiteResource) RenderError(w http.ResponseWriter, req *http.Request, status int, message string) {
	http.Redirect(w, req, "/erro", http.StatusFound)
}

func (resource *SiteResource) HandleIndex_v0(w http.ResponseWriter, req *http.Request) {
	resource.Renderer.HTML(w, http.StatusOK, "site/home", "Welcome")
}

func (resource *SiteResource) HandleError_v0(w http.ResponseWriter, req *http.Request) {
	resource.Renderer.HTML(w, http.StatusOK, "site/error", "Xii")
}