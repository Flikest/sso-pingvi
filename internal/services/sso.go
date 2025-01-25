package services

import (
	"database/sql"

	"github.com/Flikest/myMicroservices/internal/entity"
	"github.com/Flikest/myMicroservices/internal/storage"
	"github.com/gofiber/fiber"
)

type Services struct {
	storage *storage.Storage
}

type Sso interface {
	InsertUser(*entity.UserEntity) sql.Result
	GetAllUser() sql.Result
	GetUserById(id string) sql.Result
	DeleteUser(id string) sql.Result
}

var storageSSo Sso = &storage.Storage{}

func NewServices(storage *storage.Storage) *Services {
	return &Services{storage: storage}
}

func (s Services) GetUserById(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	result := storageSSo.GetUserById(id)
	ctx.JSON(result)
}

func (s Services) GetAllUser(ctx *fiber.Ctx) {
	users := storageSSo.GetAllUser()
	ctx.JSON(users)
}

func (s Services) InsertUser(ctx *fiber.Ctx) {
	body := ctx.BodyParser(entity.UserEntity{})
	result := storageSSo.InsertUser(body)
	ctx.JSON(result)
}

func (s Services) DeleteUser() {}
