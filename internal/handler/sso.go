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
	ssoRouter.Post("/logup", h.services.InsertUser)
	ssoRouter.Post("/login", h.services.LogIn)

	userRouter := router.Group("/user")
	userRouter.Get("/", h.services.GetAllUser)
	ssoRouter.Get("/:id", h.services.GetUserById)
	ssoRouter.Delete("/:id", h.services.DeleteUser)

	return *router
}
