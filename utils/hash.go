package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPwd), err
}

func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err == nil
}
