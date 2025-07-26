package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	var migrationsPath string
	var workEnvironment string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&workEnvironment, "env", "local", "working environment")

	flag.Parse()

	if err := godotenv.Load(workEnvironment + ".env"); err != nil {
		panic(fmt.Sprintf("[ MIGRATIONS ] error with loadinc .env file : %s", err))
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	migratinos, err := migrate.New(
		"file://"+migrationsPath,
		os.Getenv("POSTGRES_CONNECTING_STRING"),
	)
	if err != nil {
		panic(fmt.Sprintf("[ MIGRATIONS ] error with creating migrations: %s", err))
	}
	defer migratinos.Close()

	if err := migratinos.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("migrations not apply")
			return
		}
		panic(fmt.Sprintf("[ MIGRATIONS ] error migrations up: %s", err))
	}

	log.Println("migrations up 🚀🚀🚀")
}
