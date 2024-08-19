package models

type LoginData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
