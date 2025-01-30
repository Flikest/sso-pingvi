package main

import (
	"context"
	"os"

	"github.com/Flikest/myMicroservices/internal/handler"
	"github.com/Flikest/myMicroservices/internal/services"
	"github.com/Flikest/myMicroservices/internal/storage"
	migrations "github.com/Flikest/myMicroservices/migration"
	postgresql "github.com/Flikest/myMicroservices/pkg/clientBD/postgresql"
	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/Flikest/myMicroservices/pkg/logger"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	errors.FailOnError(err, "error creating database")

	migrations.CreateMigrations(db, "file://migration/sql")

	storage := storage.InitStorage(db, context.Background())
	services := services.NewServices(storage)
	handler := handler.InitRouter(services)
	router := handler.NewRouter()

	err = router.Listen(":3000")
	errors.FailOnError(err, "error when starting the application")

}
