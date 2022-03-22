package models

type SocialMedia struct {
	ID             int    `json:"id" example:"1"`
	Name           string `json:"name" example:"JonasP"`
	SocialMediaUrl string `json:"social_media_url" example:"www.myinsta/JonasP"`
	UserID         string `json:"user_id" example:"1"`
}
