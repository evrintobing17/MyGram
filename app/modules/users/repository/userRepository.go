package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/users"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &repo{
		db: db,
	}
}

//Add new user to db
func (r *repo) Insert(user models.User) (*models.User, error) {
	db := r.db.Create(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

//Delete existing user
func (r *repo) Delete(userId int) error {
	user := models.User{ID: userId}
	db := r.db.Delete(&user, "id = ?", userId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Get user data by email
func (r *repo) GetByEmail(email string) (*models.User, error) {
	var user models.User

	db := r.db.First(&user, "email = ?", email)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

//Get user data by email
func (r *repo) GetByID(id int) (*models.User, error) {
	var user models.User

	db := r.db.First(&user, "id = ?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}


func (r *repo) UpdatePartial(updateData map[string]interface{}) (*models.User, error) {
	id := updateData["id"]
	if id == nil {
		return nil, errors.New("field if cannot be empty")
	}
	idString := fmt.Sprintf("%v", id)
	driverID, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	var existingUser models.User
	db := r.db.First(&existingUser, "id=?", driverID)
	if db.Error != nil {
		return nil, db.Error
	}

	db = r.db.Debug().Model(&existingUser).Updates(updateData)
	if db.Error != nil {
		return nil, db.Error
	}

	return &existingUser, nil
}

func (r *repo) ExistByUsername(username string) (bool, error) {
	var user models.User
	if r.db.First(&user, "username=?", username).RecordNotFound() {
		return false, nil
	}
	return true, nil

}

func (r *repo) ExistByEmail(email string) (bool, error) {
	var user models.User
	if r.db.First(&user, "email=?", email).RecordNotFound() {
		return false, nil
	}
	return true, nil
}
