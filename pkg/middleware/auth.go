package middleware

import (
	"errors"
	"os"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized(ctx *fiber.Ctx) {
	accessToken := ctx.Get("JwtAccessPingui")
	refreshToken := ctx.Get("JwtRefreshPingui")

	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method (access)")
		}
		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})

	if err != nil {
		return
	}

	_, err = jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method (refresh)")
		}
		return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
	})
	if err != nil {
		return
	}

	ctx.Next()
}
