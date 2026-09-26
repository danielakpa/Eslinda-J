package customer

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(plainPassword string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(plainPassword))
	return err == nil
}

// CreateCustomer saves a new customer account into the database
func CreateCustomer(conn *sql.DB, name, email, hashedPassword string) error {
	_, err := conn.Exec(
		"INSERT INTO customers (name, email, password_hash) VALUES ($1, $2, $3)",
		name, email, hashedPassword,
	)
	return err
}

// GetCustomerByEmail fetches one customer's record using their email
func GetCustomerByEmail(conn *sql.DB, email string) (id int, name string, passwordHash string, err error) {
	row := conn.QueryRow("SELECT id, name, password_hash FROM customers WHERE email = $1", email)
	err = row.Scan(&id, &name, &passwordHash)
	return
}
