package easy

import "golang.org/x/crypto/bcrypt"

// BcryptPassword generates a password using bcrypt.GenerateFromPassword.
func BcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return ByteToString(bytes), nil
}

// BcryptCompare checks passwords using bcrypt.CompareHashAndPassword.
func BcryptCompare(hashed string, password string) bool {
	return bcrypt.CompareHashAndPassword(StringToByte(hashed), StringToByte(password)) == nil
}
