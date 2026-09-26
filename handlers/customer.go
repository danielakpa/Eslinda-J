package handlers

import (
	"net/http"

	"eslinda-j/customer"
)

func CustomerSignup(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedPassword, err := customer.HashPassword(password)
	if err != nil {
		http.Error(w, "Could not process password", http.StatusInternalServerError)
		return
	}

	err = customer.CreateCustomer(dbConn, name, email, hashedPassword)
	if err != nil {
		http.Error(w, "Could not create account (email may already be used)", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Account created successfully"))
}

func CustomerLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	id, _, storedHash, err := customer.GetCustomerByEmail(dbConn, email)
	if err != nil {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	if !customer.CheckPassword(password, storedHash) {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	token := customer.CreateSession(id)

	http.SetCookie(w, &http.Cookie{
		Name:  "customer_session",
		Value: token,
		Path:  "/",
	})

	w.Write([]byte("Login successful"))
}
