package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	chitracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi"
)

// Route defines a valid endpoint with the type of action supported on it
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// NewRouter returns a router handle loaded with all the supported routes
func NewRouter(routes []Route) *chi.Mux {
	r := chi.NewRouter()
	r.Use(chitracer.Middleware(chitracer.WithServiceName("dd-agent")))
	//  cors support
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	for _, route := range routes {
		r.Method(route.Method, route.Path, route.HandlerFunc)
		log.Printf("Route added: %#v\n", route)
	}

	return r
}
