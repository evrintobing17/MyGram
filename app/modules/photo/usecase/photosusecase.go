package usecase

import (
	"time"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/photo"
)

type GetPhotoResp struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      `json:"User"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UC struct {
	repo photo.PhotosRepository
}

func NewPhotoUsecase(repo photo.PhotosRepository) photo.PhotoUsecase {
	return &UC{
		repo: repo,
	}
}

func (uc *UC) AddPhoto(title, caption, url string, userId int) (*models.Photos, error) {
	modelsPhoto := models.Photos{
		Title:    title,
		Caption:  caption,
		PhotoUrl: url,
		UserID:   userId,
	}

	data, err := uc.repo.Insert(&modelsPhoto)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *UC) GetPhoto(userID int, username, email string) (interface{}, error) {
	var resp []GetPhotoResp
	data, err := uc.repo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	for _, photo := range *data {
		resp = append(resp, GetPhotoResp{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    userID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: User{
				Username: username,
				Email:    email,
			},
		})

	}
	return resp, nil
}

func (uc *UC) DeleteUserByID(photoId, userId int) error {

	err := uc.repo.CheckIfUserIDExists(photoId, userId)
	if err != nil {
		return err
	}

	errDelete := uc.repo.Delete(photoId)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func (uc *UC) UpdatePhoto(updateData map[string]interface{}) (*models.Photos, error) {

	userId := updateData["user_id"].(int)
	photoId := updateData["id"].(int)

	err := uc.repo.CheckIfUserIDExists(photoId, userId)
	if err != nil {
		return nil, err
	}

	updateData["updated_at"] = time.Now()
	userData, err := uc.repo.UpdatePartial(updateData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
