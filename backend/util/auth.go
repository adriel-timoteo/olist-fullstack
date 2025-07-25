package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/constant"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func CompareHashPassword(pwd string, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)) == nil
}

func GenerateJWTToken(userId int) (string, error) {
	secret := os.Getenv(constant.JWTSecret)
	now := time.Now()

	claims := jwt.RegisteredClaims{
		Issuer:    os.Getenv(constant.AppName),
		Subject:   strconv.Itoa(userId),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWTToken(tokenInp string) (*jwt.RegisteredClaims, error) {
	secret := os.Getenv(constant.JWTSecret)

	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenInp, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired")
		}
		return nil, err
	}

	parsedClaims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return parsedClaims, nil
}
