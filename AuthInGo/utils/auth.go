package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	// Generate a hashed version of the password using bcrypt
	//the first argument is the password to be hashed, and the second argument is the cost (or complexity) of the hashing process. A higher cost means more computational work, which can make it more secure but also slower. The default cost is 10, which is a good balance between security and performance.

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// If there was an error during hashing, return an empty string and the error
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil

}


func CheckPasswordHash(password, hash string) bool {
	// Compare the provided password with the hashed password

	// bcrypt.CompareHashAndPassword returns nil if the password matches the hash, and an error otherwise

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}