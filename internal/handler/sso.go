package handler

import (
	"github.com/Flikest/myMicroservices/internal/services"
	"github.com/gofiber/fiber"
)

type Handler struct {
	services *services.Services
}

func InitRouter(service *services.Services) *Handler {
	return &Handler{services: service}
}

func (h Handler) NewRouter() fiber.App {
	router := fiber.New()

	ssoRouter := router.Group("/sso")
	ssoRouter.Get("/user", h.services.GetAllUser)
	ssoRouter.Get("/user/:id", h.services.GetUserById)
	ssoRouter.Post("/user", h.services.InsertUser)
	ssoRouter.Delete("/user/:id", h.services.DeleteUser)

	return *router
}
