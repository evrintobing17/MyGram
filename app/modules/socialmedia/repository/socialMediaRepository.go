package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/socialmedia"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewsocialMediaRepository(db *gorm.DB) socialmedia.SocialMediaRepository {
	return &repo{
		db: db,
	}
}

//Add new SocialMedia to db
func (r *repo) Insert(SocialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	db := r.db.Create(&SocialMedia)
	if db.Error != nil {
		return nil, db.Error
	}
	return SocialMedia, nil
}

//Delete existing SocialMedia
func (r *repo) Delete(socialMediaId int) error {
	SocialMedia := models.SocialMedia{ID: socialMediaId}
	db := r.db.Delete(&SocialMedia, "id = ?", socialMediaId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Get SocialMedia data by email
func (r *repo) GetByUserID(socialMediaID int) (*[]models.SocialMedia, error) {
	var SocialMedia []models.SocialMedia

	db := r.db.Find(&SocialMedia, "user_id = ?", socialMediaID)
	if db.Error != nil {
		return nil, db.Error
	}
	return &SocialMedia, nil
}

func (r *repo) UpdatePartial(updateData map[string]interface{}) (*models.SocialMedia, error) {
	id := updateData["id"]
	if id == nil {
		return nil, errors.New("field if cannot be empty")
	}
	idString := fmt.Sprintf("%v", id)
	driverID, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	var existingsocialMedia models.SocialMedia
	db := r.db.First(&existingsocialMedia, "id=?", driverID)
	if db.Error != nil {
		return nil, db.Error
	}

	db = r.db.Debug().Model(&existingsocialMedia).Updates(updateData)
	if db.Error != nil {
		return nil, db.Error
	}

	return &existingsocialMedia, nil
}

func (r *repo) ExistBysocialMedianame(socialMedianame string) (bool, error) {
	var SocialMedia models.SocialMedia
	if r.db.First(&SocialMedia, "socialMedianame=?", socialMedianame).RecordNotFound() {
		return false, nil
	}
	return true, nil

}

func (r *repo) ExistByEmail(email string) (bool, error) {
	var SocialMedia models.SocialMedia
	if r.db.First(&SocialMedia, "email=?", email).RecordNotFound() {
		return false, nil
	}
	return true, nil
}

func (r *repo) CheckIfUserIDExists(photoId, userID int) error {
	var socialMedia models.SocialMedia
	db := r.db.Find(&socialMedia, "id = ? and user_id =?", photoId, userID)
	if db.Error != nil {
		return gorm.ErrRecordNotFound
	}
	return nil
}
