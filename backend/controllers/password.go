package controllers

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"os"
)

// HashPassword takes a plain text password and returns a salted hash
func HashPassword(password string) (string, string) {
	// Generate a random salt
	salt := make([]byte, 16)
	rand.Read(salt)

	// Hash the password and salt using SHA-256
	hash := sha256.Sum256([]byte(string(salt) + password + os.Getenv("PEPPER")))

	// Encode the salt and hash as base64 strings
	saltStr := base64.StdEncoding.EncodeToString(salt)
	hashStr := base64.StdEncoding.EncodeToString(hash[:])

	// Return the salt and salted hash
	return saltStr, hashStr
}

func VerifyPassword(password string, salt string, saltedHash string) bool {
	// Decode the salt and hash from base64 strings
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false
	}

	hashBytes, err := base64.StdEncoding.DecodeString(saltedHash)
	if err != nil {
		return false
	}

	// Hash the password and salt using SHA-256
	hash := sha256.Sum256([]byte(string(saltBytes) + password + os.Getenv("PEPPER")))

	// Compare the hashes
	return bytes.Equal(hash[:], hashBytes)
}
