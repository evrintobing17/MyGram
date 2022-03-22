package models

type Photos struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"JPG00120.jpg"`
	Caption  string `json:"caption" example:"sunday fun day"`
	PhotoUrl string `json:"photo_url" example:"https://media.istockphoto.com/photos/sunday-funday-alphabet-letter-on-wooden-background-picture-id1257920820?s=612x612"`
	UserID   string `json:"user_id" example:"1"`
	DateAudit
}
