package main

import (
	"net/http"

	_ "github.com/Haden0101/SwaggerGoProyect/docs"
	"github.com/Haden0101/SwaggerGoProyect/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title SwaggerGoProyect API
// @version 1.0
// @description API de ejemplo con Swagger en Go
// @host localhost:8080
// @BasePath /
func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Endpoint Swagger
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Endpoint GetUser
	http.HandleFunc("/user", handlers.GetUser)

	http.ListenAndServe(":8080", nil)
}
