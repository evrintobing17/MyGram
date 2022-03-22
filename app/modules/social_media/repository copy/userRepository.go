package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/evrintobing17/MyGram/app/models"
	socialmedia "github.com/evrintobing17/MyGram/app/modules/social_media"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) socialmedia.SocialMediaRepository {
	return &repo{
		db: db,
	}
}

//Add new socialMedia to db
func (r *repo) Insert(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	db := r.db.Create(&socialMedia)
	if db.Error != nil {
		return nil, db.Error
	}
	return socialMedia, nil
}

//Delete existing socialMedia
func (r *repo) Delete(socialMediaId int) error {
	socialMedia := models.SocialMedia{ID: socialMediaId}
	db := r.db.Delete(&socialMedia, "id = ?", socialMediaId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Get socialMedia data by email
func (r *repo) GetByID(socialMediaId int) (*models.SocialMedia, error) {
	var socialMedia *models.SocialMedia

	db := r.db.First(&socialMedia, "id = ?", socialMediaId)
	if db.Error != nil {
		return nil, db.Error
	}
	return socialMedia, nil
}

// func (r *repo) Update(socialMedia *models.SocialMedia) (*models.SocialMedia, error){

// }

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

	var existingSocialMedia models.SocialMedia
	db := r.db.First(&existingSocialMedia, "id=?", driverID)
	if db.Error != nil {
		return nil, db.Error
	}

	db = r.db.Debug().Model(&existingSocialMedia).Updates(updateData)
	if db.Error != nil {
		return nil, db.Error
	}

	return &existingSocialMedia, nil
}

func (r *repo) ExistBySocialMedianame(socialMedianame string) (bool, error) {
	var socialMedia models.SocialMedia
	if r.db.First(&socialMedia, "socialMedianame=?", socialMedianame).RecordNotFound() {
		return false, nil
	}
	return true, nil

}

func (r *repo) ExistByEmail(email string) (bool, error) {
	var socialMedia models.SocialMedia
	if r.db.First(&socialMedia, "email=?", email).RecordNotFound() {
		return false, nil
	}
	return true, nil
}
