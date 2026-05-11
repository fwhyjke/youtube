package server

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type APIRouter struct {
	*http.ServeMux
	routes []Route
}

func NewAPIRouter() *APIRouter {
	return &APIRouter{
		ServeMux: http.NewServeMux(),
	}
}

func (r *APIRouter) AddRoutes(routes ...Route) {
	r.routes = append(r.routes, routes...)
}

func (r *APIRouter) Handlers() map[string]http.Handler {
	handlers := make(map[string]http.Handler, len(r.routes))

	for _, route := range r.routes {
		pattern := route.Method + " /api/" + route.Path
		handlers[pattern] = route.Handler
	}

	return handlers
}
