package user

import "github.com/evrintobing17/MyGram/app/models"

type UserRepository interface {
	Insert(User *models.User) (*models.User, error)
	GetByEmail(email string) *models.User
	Update(user *models.User) (*models.User, error)
	UpdatePartial(updateData map[string]interface{}) (*models.User, error)
	Delete(userId int) error
}
