package entities

type User struct {
	Id          int    `json:"id"`
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

type UserRegister struct {
	Username string
	Password string
}

type Auth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ChangePass struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
