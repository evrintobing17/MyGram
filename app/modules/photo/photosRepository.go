package photo

import "github.com/evrintobing17/MyGram/app/models"

type PhotosRepository interface {
	Insert(Photos *models.Photos) (*models.Photos, error)
	Delete(PhotosId int) error
	GetByUserID(photoId int) (*models.Photos, error)
	GetAllByUserID(userId int) (*[]models.Photos, error)
	UpdatePartial(updateData map[string]interface{}) (*models.Photos, error)
	CheckIfUserIDExists(photoid, userID int) error
}
