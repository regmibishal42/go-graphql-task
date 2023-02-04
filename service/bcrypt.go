package service

import "golang.org/x/crypto/bcrypt"

func HashPassword(s string) string {
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	// bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashedpassword)
}

func ComparePassword(hashed string, normal string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
}
