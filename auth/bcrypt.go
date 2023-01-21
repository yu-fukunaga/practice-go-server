package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashed, _  := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func ComparePassword(hashed string, normal string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(normal))
}