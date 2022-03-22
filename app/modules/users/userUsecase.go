package users

import "github.com/evrintobing17/MyGram/app/models"

type UserUsecase interface {
	Login(email, password string) (driver *models.User, token string, err error)
	Register(username, email, password string, age int) (driver *models.User, token string, err error)
	RefreshAccessJWT(userID int) (newAccessJWT string, err error)
}
