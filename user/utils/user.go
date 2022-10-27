package utils

import "time"

type RegisterUser struct {
	Age      uint   `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisteredUser struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Username string `json:"username"`
}

type ResponseDataRegisteredUser struct {
	Status string         `json:"status"`
	Data   RegisteredUser `json:"data"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggedinUser struct {
	Token string `json:"token"`
}

type ResponseDataLoggedinUser struct {
	Status string       `json:"status"`
	Data   LoggedinUser `json:"data"`
}

type UpdateUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatedUser struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       uint       `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponseDataUpdatedUser struct {
	Status string      `json:"status"`
	Data   UpdatedUser `json:"data"`
}

type ResponseMessageDeletedUser struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
