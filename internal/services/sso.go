package services

import (
	"os"

	"github.com/Flikest/myMicroservices/internal/entity"
	"github.com/Flikest/myMicroservices/internal/storage"
	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/Flikest/myMicroservices/pkg/jwt"
	"github.com/gofiber/fiber"
)

type Services struct {
	storage *storage.Storage
}

type loginResponse struct {
	accessToken  string
	refreshToken string
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
	var body entity.UsersLogIn
	ctx.BodyParser(&body)

	id, err := s.storage.LogIn(body.Name, body.Pass)
	errors.FailOnError(err, "You introduced incorrect data")

	at, err := jwt.CreateRefreshToken(id, os.Getenv("ACCESS_SECRET_KEY"), 24)
	rt, err := jwt.CreateRefreshToken(id, os.Getenv("REFRESH_EVRET_KEY"), 1)
	errors.FailOnError(err, "JWT generation error: ")

	ctx.JSON(loginResponse{
		accessToken:  at,
		refreshToken: rt,
	})
}

func (s Services) DeleteUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	result := s.storage.DeleteUser(id)

	ctx.JSON(result)
}
