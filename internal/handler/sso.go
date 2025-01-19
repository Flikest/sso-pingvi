package handler

import (
	"github.com/Flikest/myMicroservices/internal/services"
	"github.com/gofiber/fiber"
)

type Handler struct {
	services *services.Services
}

func initRouter(service *services.Services) *Handler {
	return &Handler{services: service}
}

func (h Handler) NewRouter() fiber.App {
	router := fiber.New()

	ssoRouter := router.Group("/sso")
	ssoRouter.Get("user", h.services.GetUserById())
}
