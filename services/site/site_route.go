package site

import (
	"go-api/domain"
)

const (
	Index    = "StartPage"
	Error	= "ErrorPage"
)
const defaultBasePath = "/"

func (resource *SiteResource) generateRoutes(basePath string) *domain.Routes {
	if basePath == "" {
		basePath = defaultBasePath
	}
	var baseRoutes = domain.Routes{
		domain.Route{
			Name:           Index,
			Method:         "GET",
			Pattern:        "/",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleIndex_v0,
			},
			ACLHandler: resource.HandleIndexACL,
		},
		domain.Route{
			Name:           Error,
			Method:         "GET",
			Pattern:        "/erro",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleError_v0,
			},
			ACLHandler: resource.HandleErrorACL,
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