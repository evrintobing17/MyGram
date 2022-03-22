package users

import "github.com/evrintobing17/MyGram/app/models"

type UserRepository interface {
	Insert(user models.User) (*models.User, error)
	Delete(userId int) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	UpdatePartial(updateData map[string]interface{}) (*models.User, error)
	ExistByUsername(username string) (bool, error)
	ExistByEmail(email string) (bool, error)
}
