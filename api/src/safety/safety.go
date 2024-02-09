package safety

import "golang.org/x/crypto/bcrypt"

// Adds a hash in a string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Check if a given string and hash are the same
func CheckPassword(hashPassword, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}
