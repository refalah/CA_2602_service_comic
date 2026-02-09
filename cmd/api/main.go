package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/refalah/CA_2602_service_comic/pkg/routes"
)

func main() {
	//Load ENV
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	baseURL := os.Getenv("BASE_URL")

	//Register Routes
	r := mux.NewRouter()
	routes.RegisterRoleRoutes(r)
	routes.RegisterCreatorRoutes(r)
	routes.RegisterComicRoutes(r)
	routes.RegisterCreativeRoutes(r)

	http.Handle("/", r)

	// Log server start
	fmt.Printf("🚀 Server starting on %s\n", baseURL)
	fmt.Println("Press Ctrl+C to stop")
	
	// Start server
	log.Fatal(http.ListenAndServe(baseURL, r))
	
}