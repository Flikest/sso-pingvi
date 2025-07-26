package main

import (
	"context"
	"os"

	"github.com/Flikest/myMicroservices/internal/handler"
	"github.com/Flikest/myMicroservices/internal/services"
	"github.com/Flikest/myMicroservices/internal/storage"
	postgresql "github.com/Flikest/myMicroservices/pkg/clientBD/postgresql"
	"github.com/Flikest/myMicroservices/pkg/logger"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	log := logger.InitLogger(os.Getenv("LVL"))
	log.Info("логер запущен!")

	db, err := postgresql.DatabaseInit(&postgresql.Config{
		Context:          context.Background(),
		ConnectingString: os.Getenv("CONNECTION_STRING"),
	})
	if err != nil {
		log.Error("error creating database")
		os.Exit(1)
	}

	storage := storage.InitStorage(storage.Storage{
		DB:  db,
		Log: log,
		Ctx: context.Background(),
	})
	services := services.NewServices(services.Services{
		Storage: storage,
		Log:     log,
	})
	handler := handler.InitRouter(services)
	router := handler.NewRouter()

	if err := router.Listen(":3000"); err != nil {
		log.Error("error when starting the application")
		os.Exit(1)
	}
}
