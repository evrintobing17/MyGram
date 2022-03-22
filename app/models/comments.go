package models

type Comments struct {
	ID      int    `json:"id" example:"1"`
	UserID  int    `json:"user_id" example:"1"`
	PhotoID int    `json:"photo_id" example:"1"`
	Message string `json:"message" example:"bagus fotonya gan"`
	DateAudit
}
