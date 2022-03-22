package models

type User struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"Nanda"`
	Email    string `json:"email" example:"your.email@example.com"`
	Password string `json:"password" example:"qwerty123"`
	Age      int `json:"age" example:"19"`
	DateAudit
}
