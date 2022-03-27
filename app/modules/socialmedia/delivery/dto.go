package delivery

import "time"

type InsertReq struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type InsertResp struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type UpdateRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type UpdateResponse struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	UpdateedAt     *time.Time `json:"updated_at"`
}

type DeleteResp struct {
	Message string `json:"message"`
}
