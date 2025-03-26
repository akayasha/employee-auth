package utils

import (
	"employee-auth/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"reflect"
	"strings"
	"time"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	NIP   string `json:"nip"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// Hashing Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Check Hashing(password)
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate JWT
func GenerateJWT(nip, name, email string, role models.Role) (string, error) {
	claims := jwt.MapClaims{}
	claims["nip"] = nip
	claims["name"] = name
	claims["role"] = string(role)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// Validate Service
func ValidateStruct(input interface{}, requiredFields ...string) string {
	missingFields := []string{}
	val := reflect.ValueOf(input)

	for _, field := range requiredFields {
		f := val.FieldByName(field)
		if !f.IsValid() || (f.Kind() == reflect.String && strings.TrimSpace(f.String()) == "") {
			missingFields = append(missingFields, field)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Sprintf("Missing or empty fields: %v", strings.Join(missingFields, ", "))
	}
	return ""
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
