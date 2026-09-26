package admin

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain password and turns it into a secure hash.
// You only ever run this ONCE, when creating the admin account.
func HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword compares a login attempt against the stored hash.
// Returns true if the password is correct, false if it's wrong.
func CheckPassword(plainPassword string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(plainPassword))
	// err is nil only when the password matches
	return err == nil
}