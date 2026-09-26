package customer

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword turns a plain password into a secure hash
func HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword compares a login attempt against the stored hash
func CheckPassword(plainPassword string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(plainPassword))
	return err == nil
}

// CreateCustomer saves a new customer account into the database
func CreateCustomer(conn *sql.DB, name, phone, hashedPassword string) error {
	_, err := conn.Exec(
		"INSERT INTO customers (name, phone, password_hash) VALUES ($1, $2, $3)",
		name, phone, hashedPassword,
	)
	return err
}

// GetCustomerByPhone fetches one customer's record using their phone number
func GetCustomerByPhone(conn *sql.DB, phone string) (id int, name string, passwordHash string, err error) {
	row := conn.QueryRow("SELECT id, name, password_hash FROM customers WHERE phone = $1", phone)
	err = row.Scan(&id, &name, &passwordHash)
	return
}
