package photo

import "github.com/evrintobing17/MyGram/app/models"

type PhotosRepository interface {
	Insert(Photos *models.Photos) (*models.Photos, error)
	Delete(PhotosId int) error
	GetByID(photoId int) (*models.Photos, error)
	UpdatePartial(updateData map[string]interface{}) (*models.Photos, error)
}
