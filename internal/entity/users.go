package entity

import "github.com/google/uuid"

type UserEntity struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Pass     string    `json:"pass"`
	Avatar   string    `json:"avatar"`
	About_me string    `json:"about_me"`
}
