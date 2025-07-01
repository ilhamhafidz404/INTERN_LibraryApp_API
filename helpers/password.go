package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword menghasilkan hash dari password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPassword membandingkan password input dengan hash dari database
func CheckPassword(password, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
    return err == nil
}