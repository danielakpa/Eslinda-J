package handlers

import (
	"encoding/json"
	"net/http"

	"eslinda-j/admin"
)

// AdminLogin checks the submitted username/password and starts a session if correct
func AdminLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if !admin.CheckLogin(username, password) {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	token := admin.CreateSession()

	// Save the token in a browser cookie so future requests are recognized
	http.SetCookie(w, &http.Cookie{
		Name:  "admin_session",
		Value: token,
		Path:  "/",
	})

	w.Write([]byte("Login successful"))
}

// AddProduct handles the admin submitting a new cake through a form
func AddProduct(w http.ResponseWriter, r *http.Request) {
	// Check the admin is actually logged in before allowing this
	cookie, err := r.Cookie("admin_session")
	if err != nil || !admin.IsLoggedIn(cookie.Value) {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}

	var newProduct struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Image       string  `json:"image"`
		Category    string  `json:"category"`
		HasVariants bool    `json:"has_variants"`
	}

	err = json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	_, err = dbConn.Exec(
		"INSERT INTO products (name, price, image, category, has_variants) VALUES ($1, $2, $3, $4, $5)",
		newProduct.Name, newProduct.Price, newProduct.Image, newProduct.Category, newProduct.HasVariants,
	)
	if err != nil {
		http.Error(w, "Could not save product", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product added successfully"))
}
