package admin

import (
	"crypto/rand"
	"encoding/hex"
)

// activeSessions keeps track of valid login tokens.
// Simple approach: a token exists in this map = that admin is logged in.
var activeSessions = map[string]bool{}

// CreateSession makes a new random token and marks it as logged in
func CreateSession() string {
	tokenBytes := make([]byte, 16)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	activeSessions[token] = true
	return token
}

// IsLoggedIn checks whether a given token is a valid active session
func IsLoggedIn(token string) bool {
	return activeSessions[token]
}
