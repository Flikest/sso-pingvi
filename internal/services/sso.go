package services

import (
	"time"

	"github.com/Flikest/myMicroservices/internal/entity"
	"github.com/Flikest/myMicroservices/internal/storage"
	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
)

type Services struct {
	storage *storage.Storage
}

func createToken(username string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "todo-app",                       // Issuer
		"aud": getRole(username),                // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})
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
	s.storage.LogIn(body.Name, body.Pass)
	token, err := jwt.New()
	errors.FailOnError(err, "JWT generation error: ")
}

func (s Services) DeleteUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	result := s.storage.DeleteUser(id)
	ctx.JSON(result)
}
