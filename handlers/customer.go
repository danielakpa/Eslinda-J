package handlers

import (
	"net/http"

	"eslinda-j/customer"
)

// CustomerSignup handles creating a new customer account
func CustomerSignup(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	password := r.FormValue("password")

	hashedPassword, err := customer.HashPassword(password)
	if err != nil {
		http.Error(w, "Could not process password", http.StatusInternalServerError)
		return
	}

	err = customer.CreateCustomer(dbConn, name, phone, hashedPassword)
	if err != nil {
		http.Error(w, "Could not create account (phone may already be used)", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Account created successfully"))
}

// CustomerLogin checks phone/password and starts a session if correct
func CustomerLogin(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	password := r.FormValue("password")

	id, _, storedHash, err := customer.GetCustomerByPhone(dbConn, phone)
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
