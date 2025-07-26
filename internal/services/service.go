package services

import (
	"log/slog"

	"github.com/Flikest/myMicroservices/internal/storage"
)

type Services struct {
	Storage *storage.Storage
	Log     *slog.Logger
}

type loginResponse struct {
	accessToken  string
	refreshToken string
}

func NewServices(s Services) *Services {
	return &Services{
		Storage: s.Storage,
		Log:     s.Log,
	}
}
