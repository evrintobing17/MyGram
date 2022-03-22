package delivery

import (
	"fmt"

	"github.com/evrintobing17/MyGram/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	"github.com/evrintobing17/MyGram/app/modules/users"
	userDTO "github.com/evrintobing17/MyGram/app/modules/users/delivery/userDto"
	userUsecase "github.com/evrintobing17/MyGram/app/modules/users/usecase"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUC users.UserUsecase
}

func NewAuthHTTPHandler(r *gin.Engine, userUC users.UserUsecase, authMiddleware authmiddleware.AuthMiddleware) {
	handlers := userHandler{
		userUC: userUC,
	}

	authorized := r.Group("/users")
	{
		authorized.POST("/login", handlers.login)
		authorized.POST("/register", handlers.register)
	}
}

func (handler *userHandler) login(c *gin.Context)

// 	var loginReq driverauthenticationdto.ReqLoginAdmin

// 	errBind := c.ShouldBind(&loginReq)
// 	if errBind != nil {

// 		validations := requestvalidationerror.GetvalidationError(errBind)

// 		if len(validations) > 0 {
// 			jsonhttpresponse.BadRequest(c, validations)
// 			return
// 		}
// 		jsonhttpresponse.BadRequest(c, "")
// 		return
// 	}

// 	user, jwt, err := handler.authUseCase.Login(loginReq.Username, loginReq.Pin)
// 	if err != nil {

// 		if err == userUsecase.ErrInvalidCredential {
// 			jsonhttpresponse.Unauthorized(c, jsonhttpresponse.NewFailedResponse(err.Error()))
// 			return
// 		}

// 		jsonhttpresponse.InternalServerError(c, jsonhttpresponse.NewFailedResponse(err.Error()))
// 		return
// 	}

// 	response := driverauthenticationdto.ResLogin{
// 		User:  user,
// 		Token: jwt,
// 	}
// 	jsonhttpresponse.OK(c, response)
// 	return
// }

func (handler *userHandler) register(c *gin.Context) {
	var registerReq userDTO.Register

	errBind := c.ShouldBind(&registerReq)
	if errBind != nil {

		validations := requestvalidationerror.GetvalidationError(errBind)
		if len(validations) > 0 {

			jsonhttpresponse.BadRequest(c, validations)
			return
		}

		jsonhttpresponse.BadRequest(c, "")
		return
	}

	user, jwt, err := handler.userUC.Register(registerReq.Username, registerReq.Email, registerReq.Password, registerReq.Age)

	if err != nil {

		if err == userUsecase.ErrInvalidCredential {
			jsonhttpresponse.Unauthorized(c, err.Error())
			return
		}

		if err == userUsecase.ErrPhoneAlreadyExist ||
			err == userUsecase.ErrEmailAlreadyExist {
			jsonhttpresponse.Conflict(c, err.Error())
		}

		jsonhttpresponse.InternalServerError(c, err.Error())
		return
	}

	fmt.Println(jwt)

	response := userDTO.ResRegister{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
	}

	jsonhttpresponse.OK(c, response)
	return
}
