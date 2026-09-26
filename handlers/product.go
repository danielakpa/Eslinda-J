package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"eslinda-j/db"
)

// dbConn holds the database connection so handlers can use it.
var dbConn *sql.DB

// SetDB lets main.go pass the database connection into this package.
func SetDB(conn *sql.DB) {
	dbConn = conn
}

// GetProducts handles requests to fetch all cakes and sends them back as JSON.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAllProducts(dbConn)
	if err != nil {
		http.Error(w, "Could not fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
