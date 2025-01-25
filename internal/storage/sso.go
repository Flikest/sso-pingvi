package storage

import (
	"context"
	"database/sql"

	"log"

	"github.com/Flikest/myMicroservices/pkg/errors"
)

type Storage struct {
	db  *sql.DB
	ctx context.Context
}

func InitStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s Storage) InsertUser(u) sql.Result {
	result, err := s.db.ExecContext(s.ctx, "INSERT INTO users (id, name, pass, avatar, about_me) VALUES uuid_generate_v4(), %s, %s, %s, %s", u.Name, u.Pass, u.Avatar, u.About_me)
	errors.FailOnError(err, "error when accessing the database:")
	return result
}

func (s Storage) GetAllUser() sql.Result {
	result, err := s.db.ExecContext(s.ctx, "SELECT * FROM users")
	errors.FailOnError(err, "error when accessing the database:")
	return result
}

func (s Storage) GetUserById(id string) sql.Result {
	result, err := s.db.ExecContext(s.ctx, "SELECT * FROM users WHERE id = %s", id)
	errors.FailOnError(err, "error when accessing the database:")
	return result
}

func (s Storage) DeleteUser(id string) sql.Result {
	result, err := s.db.ExecContext(s.ctx, "DELETE FROM users WHERE id=%s", id)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
