package dto

type ReqRegisterUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id    int    `json:"user_id"`
	Email string `json:"email"`
}

type ReqLoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
}
