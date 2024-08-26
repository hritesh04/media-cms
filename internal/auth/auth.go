package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/helper"
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

func IsAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if role := ctx.GetHeader("role"); role == "AUTHOR" {
			fmt.Println(ctx.Request.Header)
			ctx.Next()
			return
		}
		helper.ReturnFailed(ctx, http.StatusUnauthorized, fmt.Errorf("user is not an author"))
		ctx.Abort()
	}
}

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("media")
		if err != nil {
			helper.ReturnFailed(ctx, http.StatusBadRequest, err)
			ctx.Abort()
			return
		}
		claims, err := ValidateUser(tokenString)
		if err != nil {
			helper.ReturnFailed(ctx, http.StatusBadRequest, err)
			ctx.Abort()
			return
		}
		fmt.Println(claims)
		if userId, ok := claims["userID"].(string); ok {
			fmt.Println("Setting user ID:", userId)
			ctx.Request.Header.Set("userID", userId)
		} else {
			helper.ReturnFailed(ctx, http.StatusBadRequest, fmt.Errorf("invalid token: no user ID"))
			ctx.Abort()
			return
		}
		if role, ok := claims["role"].(string); ok {
			fmt.Println("ROLE")
			ctx.Request.Header.Set("role", role)
		}
		fmt.Println(ctx.Request.Header)
		ctx.Next()
	}
}

func ValidateUser(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return jwt.MapClaims{}, nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return jwt.MapClaims{}, fmt.Errorf("token is invalid")
}

func GenerateToken(id uint, role domain.Role) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.FormatUint(uint64(id), 10),
		"role":   role.Value(),
		"exp":    time.Now().Add(time.Hour).Unix(),
		"issue":  time.Now().Unix(),
	})
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

func ComparePassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
