package todo

import (
	"go-api/domain"
)

const (
	List      = "ListTodos"
	Create      = "CreateTodos"
)
const defaultBasePath = "/api/todo"

func (resource *TodoResource) generateRoutes(basePath string) *domain.Routes {
	if basePath == "" {
		basePath = defaultBasePath
	}
	var baseRoutes = domain.Routes{
		domain.Route{
			Name:           List,
			Method:         "GET",
			Pattern:        "/todo",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleIndexTodo_v0,
			},
			ACLHandler: resource.HandleTodoListACL,
		},
		domain.Route{
			Name:           List,
			Method:         "GET",
			Pattern:        basePath,
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleListTodo_v0,
			},
			ACLHandler: resource.HandleTodoListACL,
		},
		domain.Route{
			Name:           Create,
			Method:         "POST",
			Pattern:        basePath,
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleCreateTodo_v0,
			},
			ACLHandler: resource.HandleTodoCreateACL,
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