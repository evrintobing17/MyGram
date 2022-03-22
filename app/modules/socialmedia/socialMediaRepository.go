package socialmedia

import "github.com/evrintobing17/MyGram/app/models"

type SocialMediaRepository interface {
	Insert(SocialMedia *models.SocialMedia) (*models.SocialMedia, error)
	Delete(SocialMediaId int) error
	GetByID(SocialMediaId int) (*models.SocialMedia, error)
	UpdatePartial(updateData map[string]interface{}) (*models.SocialMedia, error)
}
