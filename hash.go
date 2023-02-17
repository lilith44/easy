package easy

import "golang.org/x/crypto/bcrypt"

func BcryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func BcryptCompare(hashed []byte, password []byte) bool {
	return bcrypt.CompareHashAndPassword(hashed, password) == nil
}
