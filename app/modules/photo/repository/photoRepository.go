package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/photo"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewPhotosRepository(db *gorm.DB) photo.PhotosRepository {
	return &repo{
		db: db,
	}
}

//Add new photo to db
func (r *repo) Insert(photo *models.Photos) (*models.Photos, error) {
	db := r.db.Create(&photo)
	if db.Error != nil {
		return nil, db.Error
	}
	return photo, nil
}

//Delete existing photo
func (r *repo) Delete(photoId int) error {
	photo := models.Photos{ID: photoId}
	db := r.db.Delete(&photo, "id = ?", photoId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Get photo data by email
func (r *repo) GetByUserID(photoId int) (*models.Photos, error) {
	var photo models.Photos
	fmt.Println(photoId)
	db := r.db.Find(&photo, "user_id = ?", photoId)
	if db.Error != nil {
		return nil, db.Error
	}
	return &photo, nil
}

func (r *repo) UpdatePartial(updateData map[string]interface{}) (*models.Photos, error) {
	id := updateData["id"]
	if id == nil {
		return nil, errors.New("field if cannot be empty")
	}
	idString := fmt.Sprintf("%v", id)
	photoID, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	var existingPhotos models.Photos
	db := r.db.First(&existingPhotos, "id=?", photoID)
	if db.Error != nil {
		return nil, db.Error
	}

	db = r.db.Debug().Model(&existingPhotos).Updates(updateData)
	if db.Error != nil {
		return nil, db.Error
	}

	return &existingPhotos, nil
}

func (r *repo) GetAllByUserID(userId int) (*[]models.Photos, error) {
	var photo []models.Photos

	db := r.db.Find(&photo, "user_id = ?", userId)
	if db.Error != nil {
		return nil, db.Error
	}
	return &photo, nil
}

func (r *repo) CheckIfUserIDExists(photoId, userID int) error {
	var photo models.Photos
	db := r.db.Find(&photo, "id = ? and user_id =?", photoId, userID)
	if db.Error != nil {
		return gorm.ErrRecordNotFound
	}
	return nil
}
