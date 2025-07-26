package services

import (
	"os"

	"github.com/Flikest/myMicroservices/internal/storage"
	"github.com/Flikest/myMicroservices/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (s *Services) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		s.Log.Warn("no id")
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON("no id")
	}

	result, err := s.Storage.GetUserById(id)
	if err != nil {
		s.Log.Error("error wiht getting user by id: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("error with getting user by id")
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(result)
}

func (s *Services) GetAllUser(ctx *fiber.Ctx) error {
	users, err := s.Storage.GetAllUser()
	if err != nil {
		s.Log.Error("error with getting all user: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("error with getting all user")
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(users)
}

func (s *Services) InsertUser(ctx *fiber.Ctx) error {
	var body storage.User
	if err := ctx.BodyParser(&body); err != nil {
		s.Log.Warn("no body", "err", err)
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON("no body")
	}

	err := s.Storage.InsertUser(body)
	if err != nil {
		s.Log.Error("error with inserting user: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("errro with created user")
	}

	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(body)
}

func (s *Services) LogIn(ctx *fiber.Ctx) error {
	var body storage.UsersLogIn
	if err := ctx.BodyParser(&body); err != nil {
		s.Log.Warn("no body", "err", err)
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON("no body")
	}

	id, err := s.Storage.LogIn(body)
	if err != nil {
		s.Log.Error("error with logining user: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("error with logining user")
	}

	accessToken, err := jwt.CreateRefreshToken(id, os.Getenv("ACCESS_SECRET_KEY"), 24)
	if err != nil {
		s.Log.Error("JWT access generation error: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("JWT access generation error")
	}

	refreshToken, err := jwt.CreateRefreshToken(id, os.Getenv("REFRESH_EVRET_KEY"), 1)
	if err != nil {
		s.Log.Error("JWT refresh generation error: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("JWT refresh generation error")
	}

	return ctx.JSON(loginResponse{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	})
}

func (s *Services) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		s.Log.Warn("no id")
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON("no id")
	}

	err := s.Storage.DeleteUser(id)
	if err != nil {
		s.Log.Error("error with deleting user: ", "err", err)
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON("error with deleting user")
	}

	ctx.Status(fiber.StatusNoContent)
	return ctx.JSON("")
}
