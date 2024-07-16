package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword hashes a password using SHA-256 algorithm
func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

// Example of other utility functions you might add:

// ValidatePassword checks if a provided password matches a hashed password
func ValidatePassword(password, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}
