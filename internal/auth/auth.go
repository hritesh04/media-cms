package auth

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var secret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	secret = []byte(os.Getenv("SECRET"))
}

func GenerateToken(id uint, role string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"role":  role,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"issue": time.Now().Unix(),
	})
	fmt.Println(string(secret))
	token, err := claims.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error generating a token")
	}
	return token, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password : %v", err)
	}
	return string(hash), nil
}

func ComparePassword(password string) {

}
