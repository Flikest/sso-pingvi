package services

import "github.com/Flikest/myMicroservices/internal/storage"

type Services struct {
	storage *storage.Storage
}

func NewServices(storage *storage.Storage) {

}

func (s Services) GetUserById() {

}

func (s Services) GetAllUser() {}

func (s Services) InsertUser() {}

func (s Services) DeleteUser() {}
