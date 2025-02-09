package services

import (
	"os"

	"github.com/Flikest/myMicroservices/internal/entity"
	"github.com/Flikest/myMicroservices/internal/storage"
	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
)

type Services struct {
	storage *storage.Storage
}

func createToken(name, pass, avatar, about_me string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"password": pass,
		"avatar":   avatar,
		"about_me": about_me,
	})

	tokenString, err := claims.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewServices(storage *storage.Storage) *Services {
	return &Services{storage: storage}
}

func (s Services) GetUserById(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	result := s.storage.GetUserById(id)

	ctx.JSON(result)
}

func (s Services) GetAllUser(ctx *fiber.Ctx) {
	users := s.storage.GetAllUser()

	ctx.JSON(users)
}

func (s Services) InsertUser(ctx *fiber.Ctx) {
	var body entity.UserEntity
	ctx.BodyParser(&body)

	result := s.storage.InsertUser(&body)

	ctx.JSON(result)
}

func (s Services) LogIn(ctx *fiber.Ctx) {
	var body entity.UserEntity
	ctx.BodyParser(&body)

	row := s.storage.LogIn(body.Name, body.Pass)

	token, err := createToken()
	errors.FailOnError(err, "JWT generation error: ")

	ctx.JSON(token)
}

func (s Services) DeleteUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	result := s.storage.DeleteUser(id)

	ctx.JSON(result)
}
