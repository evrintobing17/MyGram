package photo

import "github.com/evrintobing17/MyGram/app/models"

type PhotoUsecase interface {
	AddPhoto(title, caption, url string, userId int) (*models.Photos, error)
	GetPhoto(userID int, username, email string) (interface{}, error)
	UpdatePhoto(updateData map[string]interface{}) (user *models.Photos, err error)
	DeleteUserByID(photoId, userId int) error
}
