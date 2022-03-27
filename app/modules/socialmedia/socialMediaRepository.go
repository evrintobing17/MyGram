package socialmedia

import "github.com/evrintobing17/MyGram/app/models"

type SocialMediaRepository interface {
	Insert(SocialMedia *models.SocialMedia) (*models.SocialMedia, error)
	Delete(SocialMediaId int) error
	GetByUserID(userId int) (*[]models.SocialMedia, error)
	UpdatePartial(updateData map[string]interface{}) (*models.SocialMedia, error)
	CheckIfUserIDExists(photoId, userID int) error
}
