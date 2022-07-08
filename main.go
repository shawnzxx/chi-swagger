package main

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
	"net/http"
	"time"

	"github.com/ganpatagarwal/chi-swagger/handlers"

	_ "github.com/ganpatagarwal/chi-swagger/docs"
	"github.com/ganpatagarwal/chi-swagger/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title chi-swagger example APIs
// @version 1.0
// @description chi-swagger example APIs
// @BasePath /
func main() {
	tracer.Start(
		tracer.WithEnv("prod"),
		tracer.WithService("dd-agent"),
		// or, for a non-default TCP connection:
		tracer.WithAgentAddr("localhost:8126"),
	)

	// When the tracer is stopped, it will flush everything it has to the Datadog Agent before quitting.
	// Make sure this line stays in your main function.
	defer tracer.Stop()

	var timeout = 2 * time.Minute

	var routes = []router.Route{
		router.Route{
			Method:      "GET",
			Path:        "/",
			HandlerFunc: handlers.RootHandler,
		},
	}

	log.Println("Launching the server")
	r := router.NewRouter(routes)
	r.Mount("/swagger", httpSwagger.WrapHandler)

	server := http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      r,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Failed to launch api server:%+v\n", err)
	}
}
