package main

import (
	"os"

	"github.com/Flikest/myMicroservices/internal/handler"
	"github.com/Flikest/myMicroservices/internal/services"
	"github.com/Flikest/myMicroservices/internal/storage"
	migrations "github.com/Flikest/myMicroservices/migration"
	postgresql "github.com/Flikest/myMicroservices/pkg/clientBD/postgresql"
	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/Flikest/myMicroservices/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := os.Getenv("LVL")
	log := logger.InitLogger(env)
	log.Debug("логер запущен!")

	db, err := postgresql.NewDatabase(&postgresql.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	errors.FailOnError(err, "error creating database")

	migrations.CreateMigrations(db, "")

	storage := storage.InitStorage(db)
	services := services.NewServices(storage)
	handler := handler.InitRouter(services)
	router := handler.NewRouter()

	router.Listen(":3000")
}
