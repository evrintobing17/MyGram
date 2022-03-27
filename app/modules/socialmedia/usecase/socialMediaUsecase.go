package usecase

import (
	"time"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/socialmedia"
)

type SocialMediaData struct {
	Data []GetSocialMediaResp `json:"social_medias"`
}

type GetSocialMediaResp struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	User           `json:"User"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UC struct {
	socialMediaRepo socialmedia.SocialMediaRepository
}

func NewsocialMediaUsecase(socialMediaRepo socialmedia.SocialMediaRepository) socialmedia.SocialMediaUsecase {
	return &UC{
		socialMediaRepo: socialMediaRepo,
	}
}

func (uc *UC) AddSocialMedia(name, url string, userId int) (*models.SocialMedia, error) {
	socialMedia := models.SocialMedia{
		Name:           name,
		SocialMediaUrl: url,
		UserID:         userId,
	}

	data, err := uc.socialMediaRepo.Insert(&socialMedia)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (uc *UC) GetSocialMedia(userId int, username, email string) (interface{}, error) {
	var resp []GetSocialMediaResp
	data, err := uc.socialMediaRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	for _, socialMedia := range *data {
		resp = append(resp, GetSocialMediaResp{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: User{
				ID:       userId,
				Username: username,
				Email:    email,
			},
		})

	}
	response := SocialMediaData{
		Data: resp,
	}
	return response, nil
}
func (uc *UC) DeleteSocialMediaByID(socialMediaId, userId int) error {
	err := uc.socialMediaRepo.CheckIfUserIDExists(socialMediaId, userId)
	if err != nil {
		return err
	}

	errDelete := uc.socialMediaRepo.Delete(socialMediaId)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
func (uc *UC) UpdateSocialMedia(updateData map[string]interface{}) (*models.SocialMedia, error) {
	userId := updateData["user_id"].(int)
	photoId := updateData["id"].(int)

	err := uc.socialMediaRepo.CheckIfUserIDExists(photoId, userId)
	if err != nil {
		return nil, err
	}

	updateData["updated_at"] = time.Now()
	userData, err := uc.socialMediaRepo.UpdatePartial(updateData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
