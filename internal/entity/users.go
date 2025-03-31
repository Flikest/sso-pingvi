package entity

type (
	UserEntity struct {
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
