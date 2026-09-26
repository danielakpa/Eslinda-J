package main

import (
	"log"
	"mime"
	"net/http"

	"eslinda-j/db"
	"eslinda-j/handlers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mime.AddExtensionType(".js", "application/javascript")

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
	// Explicitly tell the server that .js files are JavaScript
	// (fixes a MIME type issue that happens on some Linux setups)
	http.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/shop.html")
	})

	http.HandleFunc("/customer-signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/customer-signup.html")
	})

	http.HandleFunc("/customer-login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/customer-login.html")
	})

	// Give the handlers package access to the database
	handlers.SetDB(database)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// New route: fetching all cakes
	http.HandleFunc("/api/products", handlers.GetProducts)
	http.HandleFunc("/admin/login", handlers.AdminLogin)
	http.HandleFunc("/admin/products/add", handlers.AddProduct)
	http.HandleFunc("/customer/signup", handlers.CustomerSignup)
	http.HandleFunc("/customer/login", handlers.CustomerLogin)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
