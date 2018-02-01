package session

import (
	"go-api/domain"
)

const (
	GetSession    = "GetSession"
	CreateSession = "CreateSession"
	DeleteSession = "DeleteSession"
)

const defaultBasePath = "/api/sessions"

func (resource *SessionResource) generateRoutes(basePath string) *domain.Routes {
	if basePath == "" {
		basePath = defaultBasePath
	}
	var baseRoutes = domain.Routes{

		domain.Route{
			Name:           GetSession,
			Method:         "GET",
			Pattern:        basePath,
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleGetSession_v0,
			},
			ACLHandler: resource.HandleGetSessionACL,
		},
		domain.Route{
			Name:           CreateSession,
			Method:         "POST",
			Pattern:        basePath,
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleCreateSession_v0,
			},
			ACLHandler: resource.HandleCreateSessionACL,
		},
		domain.Route{
			Name:           DeleteSession,
			Method:         "DELETE",
			Pattern:        basePath,
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleDeleteSession_v0,
			},
			ACLHandler: resource.HandleDeleteSessionACL,
		},
	}

	routes := domain.Routes{}

	for _, route := range baseRoutes {
		r := domain.Route{
			Name:           route.Name,
			Method:         route.Method,
			Pattern:        route.Pattern,
			DefaultVersion: route.DefaultVersion,
			RouteHandlers:  route.RouteHandlers,
			ACLHandler:     route.ACLHandler,
		}
		routes = routes.Append(&domain.Routes{r})
	}
	resource.routes = &routes
	return resource.routes
}