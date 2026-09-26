package main

import (
	"log"
	"net/http"

	"eslinda-j/db"
	"eslinda-j/handlers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file:", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	database := db.Connect(dbHost, dbPort, dbUser, dbPassword, dbName)
	defer database.Close()

	// Give the handlers package access to the database
	handlers.SetDB(database)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// New route: fetching all cakes
	http.HandleFunc("/api/products", handlers.GetProducts)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
