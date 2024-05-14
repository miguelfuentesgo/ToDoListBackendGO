package router

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type RouteHandlers struct {
	router *fasthttprouter.Router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fasthttp.RequestHandler
}

func NewRouter() *fasthttprouter.Router {
	handlers := &RouteHandlers{
		router: fasthttprouter.New(),
	}
	handlers.addRoutes()
	return handlers.router
}

func (r *RouteHandlers) addRoutes() {
	log.Printf("Registering routes - start")
	defer log.Printf("Registering routes - completed")

	for _, route := range r.routes() {
		var handler fasthttp.RequestHandler
		handler = route.HandlerFunc
		switch route.Method {
		case "GET":
			r.router.GET(route.Pattern, handler)
		case "POST":
			r.router.POST(route.Pattern, handler)
		case "PUT":
			r.router.PUT(route.Pattern, handler)
		case "DELETE":
			r.router.DELETE(route.Pattern, handler)
		case "PATCH":
			r.router.PATCH(route.Pattern, handler)
		}
	}

}

func (r *RouteHandlers) routes() []Route {
	routes := make([]Route, 0)
	routes = append(routes, r.routesToDoApp()...)
	return routes
}
