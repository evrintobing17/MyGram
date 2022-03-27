package delivery

import "time"

type InsertReq struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id"`
}

type InsertResp struct {
	ID        int        `json:"id"`
	Message   string     `json:"message" binding:"required"`
	PhotoID   int        `json:"photo_id"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdateRequest struct {
	Message string `json:"message" binding:"required"`
}

type UpdateResponse struct {
	ID        int        `json:"id"`
	Message   string     `json:"message"`
	PhotoID   int        `json:"photo_id"`
	UserID    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"update_at"`
}

type DeleteResp struct {
	Message string `json:"message"`
}
