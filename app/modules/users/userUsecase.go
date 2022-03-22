package users

import (
	"github.com/evrintobing17/MyGram/app/models"
)

type UserUsecase interface {
	Login(email, password string) (user *models.User, token string, err error)
	Register(username, email, password string, age int) (user *models.User, token string, err error)
	RefreshAccessJWT(userID int) (newAccessJWT string, err error)
	DeleteUserByID(userId int) error
	UpdateUser(updateData map[string]interface{}) (user *models.User, err error)
}
