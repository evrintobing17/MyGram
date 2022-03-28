package socialmedia

import (
	"github.com/evrintobing17/MyGram/app/models"
)

type SocialMediaUsecase interface {
	AddSocialMedia(name, url string, userId int) (*models.SocialMedia, error)
	GetSocialMedia(userId int, username string) (interface{}, error)
	DeleteSocialMediaByID(socialMediaId, userId int) error
	UpdateSocialMedia(updateData map[string]interface{}) (*models.SocialMedia, error)
}
