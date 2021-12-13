package security

import "golang.org/x/crypto/bcrypt"

//Hash the password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

//VerifyPassword verifies the password and the hash
func VerifyPassword(passwordWithHash, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(stringPassword))
}
