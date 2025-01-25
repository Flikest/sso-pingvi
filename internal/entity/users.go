package entity

type UserEntity struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Pass     string `json:"pass"`
	Avatar   string `json:"avatar"`
	About_me string `json:"about_me"`
}
