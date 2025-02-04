package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Flikest/myMicroservices/internal/entity"
	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/Flikest/myMicroservices/rabbitmq"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Storage struct {
	db  *sql.DB
	ctx context.Context
}

func InitStorage(db *sql.DB, ctx context.Context) *Storage {
	return &Storage{db: db, ctx: ctx}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s Storage) InsertUser(u *entity.UserEntity) sql.Result {
	password, err := HashPassword(u.Pass)
	errors.FailOnError(err, "password hash error: ")

	id := uuid.New()

	result, err := s.db.ExecContext(s.ctx, "INSERT INTO users (id, name, pass, avatar, about_me) VALUES ($1, $2, $3, $4, $5) RETURNING *", id, u.Name, password, u.Avatar, u.About_me)
	errors.FailOnError(err, "error when accessing the database:")

	queryExchenge := fmt.Sprintf("%s,%s,%s,%s,%s", id, u.Name, password, u.Avatar, u.About_me)
	rabbitmq.Send(queryExchenge)

	return result
}

func (s Storage) LogIn(name string, password string) *sql.Row {
	result := s.db.QueryRowContext(s.ctx, "SELECT * FROM users WHERE name=$1, pass=$2", name, password)
	return result
}

func (s Storage) GetAllUser() []entity.UserEntity {
	rows, err := s.db.QueryContext(s.ctx, "SELECT * FROM users")
	errors.FailOnError(err, "error when accessing the database:")
	result := []entity.UserEntity{}
	for rows.Next() {
		var users = entity.UserEntity{}
		if err := rows.Scan(&users.Id, &users.Name, &users.Pass, &users.Avatar, &users.About_me); err != nil {
			log.Fatalln(err)
		}
		result = append(result, users)
	}
	return result
}

func (s Storage) GetUserById(id string) *sql.Row {
	result := s.db.QueryRowContext(s.ctx, "SELECT * FROM users WHERE id=$1", id)
	return result
}

func (s Storage) DeleteUser(id string) sql.Result {
	result, err := s.db.ExecContext(s.ctx, "DELETE FROM users WHERE id=$1", id)
	errors.FailOnError(err, "error when accessing the database:")
	return result
}
