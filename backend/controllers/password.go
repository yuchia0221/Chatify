package controllers

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// HashPassword takes a plain text password and returns a salted hash
func HashPassword(password string) (string, string, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", "", err
	}

	// Hash the password and salt using SHA-256
	hash := sha256.Sum256([]byte(string(salt) + password + os.Getenv("PEPPER")))

	// Encode the salt and hash as base64 strings
	saltStr := base64.StdEncoding.EncodeToString(salt)
	hashStr := base64.StdEncoding.EncodeToString(hash[:])

	// Return the salt and salted hash
	return saltStr, hashStr, nil
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

func GenerateToken(username string) (string, error) {
	// Set the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	// Create the claims for the token
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
