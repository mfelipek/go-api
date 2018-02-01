package server

import (
	"log"
	"fmt"
	"errors"	
	"context"
	"net/http"
	"go-api/domain"
	"github.com/gorilla/mux"
)

// Router type
type Router struct {
	*mux.Router
	ac  domain.IAccessController
	ctx context.Context
}

// matcherFunc matches the handler to the correct API version based on its `accept` header
// TODO: refactor matcher function as server.Config
func matcherFunc(r domain.Route, defaultHandler http.HandlerFunc, ctx context.Context, ac domain.IAccessController) func(r *http.Request, rm *mux.RouteMatch) bool {
	return func(req *http.Request, rm *mux.RouteMatch) bool {
	
		// make sure routes under the same pattern but with diff method are not matched
		if req.Method != "" && r.Method != req.Method {
			return false;
		}
		
		acceptHeaders := domain.NewAcceptHeadersFromString(req.Header.Get("accept"))
		foundHandler := defaultHandler
		
		// try to match a handler to the specified `version` params
		// else we will fall back to the default handler
		for _, h := range acceptHeaders {
			m := h.MediaType
			
			// check if media type is `application/json` type or `application/[*]+json` suffix
			if !(m.Type == "application" && (m.SubType == "json" || m.Suffix == "json")) {
				continue
			}

			// if its the right application type, check if a version specified
			version, hasVersion := m.Parameters["version"]
			if !hasVersion {
				continue
			}
			if handler, ok := r.RouteHandlers[domain.RouteHandlerVersion(version)]; ok {
				// found handler for specified version
				foundHandler = handler
				break
			}
		}
		
		if ac != nil {
			log.Println("NewContextHandler ", r.Name)
			rm.Handler = ac.NewContextHandler(r.Name, foundHandler)
		} else {
			rm.Handler = foundHandler
		}
		return true
	}
}

// NewRouter Returns a new Router object
func NewRouter(ctx context.Context, ac domain.IAccessController, staticContentPath string) *Router {
	router := mux.NewRouter().StrictSlash(true)

	ServeStatic(router, staticContentPath)

	return &Router{router, ac, ctx}
}

func (router *Router) AddRoutes(routes *domain.Routes) *Router {
	if routes == nil {
		return router
	}
	for _, route := range *routes {

		// get the defaultHandler for current route at init time so that we can safely panic
		// if it was not defined
		defaultHandler, ok := route.RouteHandlers[route.DefaultVersion]
		if !ok {
			// server/router instantiation error
			// its safe to throw panic here
			panic(errors.New(fmt.Sprintf("Routes definition error, missing default route handler for version `%v` in `%v`",
				route.DefaultVersion, route.Name)))
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			MatcherFunc(matcherFunc(route, defaultHandler, router.ctx, router.ac))
		
		if router.ac != nil {
			router.ac.AddHandler(route.Name, route.ACLHandler)
		}
			
	}
	return router
}

func (router *Router) AddResources(resources ...domain.IResource) *Router {
	for _, resource := range resources {
		if resource.Routes() == nil {
			// server/router instantiation error
			// its safe to throw panic here
			panic(errors.New(fmt.Sprintf("Routes definition missing: %v", resource)))
		}
		router.AddRoutes(resource.Routes())
	}
	return router
}

func ServeStatic(router *mux.Router, staticDirectory string) {
    staticPaths := map[string]string{
        "styles":           staticDirectory + "/styles/",
        "bower_components": staticDirectory + "/bower_components/",
        "images":           staticDirectory + "/images/",
        "scripts":          staticDirectory + "/scripts/",
    }
    for pathName, pathValue := range staticPaths {
        pathPrefix := "/" + pathName + "/"
        router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
            http.FileServer(http.Dir(pathValue))))
    }
}