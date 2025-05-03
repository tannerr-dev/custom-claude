package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"database/sql"
	"github.com/joho/godotenv"
)


func main() {
	// Log Initializer
	logInstance := initializeLogger()

	// Load .env file, this mergest .env into the os env variables
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// // Database connection
	// dbConnStr := os.Getenv("DATABASE_URL")
	// if dbConnStr == "" {
	// 	log.Fatalf("DATABASE_URL not set in environment")
	// }
	// db, err := sql.Open("postgres", dbConnStr)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()
	//
	// // Initialize Data Repository for Movies
	// movieRepo, err := data.NewMovieRepository(db, logInstance)
	// if err != nil {
	// 	log.Fatalf(("Failed to initialize repository"))
	// }
	//

	// http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)

	catchAllClientRoutesHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r, "./public/index.html")
	}

	// Handler for static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving the files")

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}
}
