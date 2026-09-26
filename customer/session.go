package customer

import (
	"crypto/rand"
	"encoding/hex"
)

// activeSessions maps a session token to the logged-in customer's ID
var activeSessions = map[string]int{}

// CreateSession makes a new token for a logged-in customer
func CreateSession(customerID int) string {
	tokenBytes := make([]byte, 16)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	activeSessions[token] = customerID
	return token
}

// GetCustomerID returns the customer ID for a session token, or 0 if not logged in
func GetCustomerID(token string) int {
	return activeSessions[token]
}
