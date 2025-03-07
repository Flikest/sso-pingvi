package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func CreateAccessToken(uid string, accesSecret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &JwtCustomClaims{
		ID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	at, err := token.SignedString([]byte(accesSecret))
	if err != nil {
		return "", err
	}

	return at, err

}

func CreateRefreshToken(uid string, refreshSecret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	claims := &JwtCustomClaims{
		ID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	rt, err := token.SignedString([]byte(refreshSecret))
	if err != nil {
		return "", err
	}

	return rt, err
}

func ExtractPyloadFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", errors.New("Invalid token")
	}

	return claims["id"].(string), nil
}
