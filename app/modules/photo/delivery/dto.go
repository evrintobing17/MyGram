package delivery

import "time"

type InsertReq struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type InsertResp struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type PhotoResp struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`
	User
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type UpdateResponse struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"update_at"`
}

type DeleteResp struct {
	Message string `json:"message"`
}
