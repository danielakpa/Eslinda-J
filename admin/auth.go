package admin

import (
	"encoding/json"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// AdminCredentials matches the structure of admin.json
type AdminCredentials struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

// HashPassword turns a plain password into a secure hash (used once, to set up admin.json)
func HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(hash), err
}

// loadCredentials reads admin.json from disk
func loadCredentials() (AdminCredentials, error) {
	var pass AdminCredentials

	fileData, err := os.ReadFile("admin/admin.json")
	if err != nil {
		return pass, err
	}

	err = json.Unmarshal(fileData, &pass)
	return pass, err
}

// CheckLogin verifies a username/password attempt against admin.json
func CheckLogin(username, password string) bool {
	pass, err := loadCredentials()
	if err != nil {
		return false
	}

	if username != pass.Username {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass.PasswordHash), []byte(password))
	return err == nil
}
