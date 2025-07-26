package storage

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Pass     string `json:"pass"`
		Avatar   string `json:"avatar"`
		About_me string `json:"about_me"`
	}

	UsersLogIn struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Storage) InsertUser(u User) error {
	password, err := HashPassword(u.Pass)
	if err != nil {
		s.Log.Error("password hash error: ")

	}

	id := uuid.New()

	_, err = s.DB.Exec(s.Ctx, "INSERT INTO users (id, name, email, pass, avatar, about_me) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", id, u.Name, u.Email, password, u.Avatar, u.About_me)
	if err != nil {
		s.Log.Error("error when accessing the database: ", "err", err)
		return err
	}

	return nil
}

func (s *Storage) LogIn(login UsersLogIn) (uuid.UUID, error) {
	var id uuid.UUID
	var pass string

	row := s.DB.QueryRow(s.Ctx, "SELECT id, pass FROM users WHERE name=$1", login.Name)
	if err := row.Scan(&id, &pass); err != nil {
		s.Log.Error("error with selecting id and password from users table by name param: ", "err", err)
		return uuid.Nil, err
	}

	if CheckPasswordHash(login.Pass, pass) {
		return id, nil
	} else {
		return uuid.Nil, fmt.Errorf("you entered incorrect data")
	}
}

func (s *Storage) GetAllUser() ([]User, error) {
	rows, err := s.DB.Query(s.Ctx, "SELECT * FROM users")
	if err != nil {
		s.Log.Error("error when accessing the database:", "err", err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user = User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Pass, &user.Avatar, &user.About_me); err != nil {
			s.Log.Error("error with scaning users table: ", "err", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Storage) GetUserById(id string) (User, error) {
	var user User
	err := s.DB.QueryRow(s.Ctx, "SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Pass, &user.Avatar, &user.About_me)
	if err != nil {
		s.Log.Error("error wiht selecting data form user table by id: ", "err", err)
		return user, err
	}

	return user, nil
}

func (s *Storage) DeleteUser(id string) error {
	_, err := s.DB.Exec(s.Ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		s.Log.Error("error when accessing the database: ", "err", err)
		return err
	}
	return nil
}
