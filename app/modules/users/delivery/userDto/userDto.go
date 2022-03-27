package userdto

import "time"

type Register struct {
	Age      int    `json:"age" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

type ResRegister struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type ReqLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

type ResLogin struct {
	Jwt string `json:"jwt"`
}

type ReqUpdate struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type RespUpdate struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       int        `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteResp struct {
	Message string `json:"message"`
}
