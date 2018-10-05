package main

import (
	"backend_go/internal/app/business"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// set default configuration
	serverPort := 9999
	serverHost := "http://localhost"

	envAppHost := os.Getenv("APPLICATION_HOST")
	if envAppHost != "" {
		serverHost = envAppHost
	}

	envAppPort := os.Getenv("APPLICATION_PORT")
	if envAppPort != "" {
		serverPort, _ = strconv.Atoi(envAppPort)
	}

	fmt.Printf("Starting application on %s\n", fmt.Sprintf("%s:%d", serverHost, serverPort))

	// Initialize router and some middleware
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.Headers("Content-Type", "application/json")

	// Register route
	router.Path("/health").HandlerFunc(business.HealthCheck).Methods("GET")
	router.Path("/regions").HandlerFunc(business.GetAllRegions).Methods("GET")
	router.Path("/verticals").HandlerFunc(business.GetAllVerticals).Methods("GET")
	router.Path("/product_lines").HandlerFunc(business.GetAllProductLines).Methods("GET")

	// Run the http server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		w.Header().Add("Content-Type", jsonapi.MediaType)
		next.ServeHTTP(w, r)
	})
}
