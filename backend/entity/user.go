package entity

type ReqRegisterUser struct {
	Email    string
	Password string
}

type User struct {
	Id       int
	Email    string
	Password string
}

type ReqLoginUser struct {
	Email    string
	Password string
}

type Token struct {
	Token string
}
