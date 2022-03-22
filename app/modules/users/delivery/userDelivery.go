package delivery

import (
	"github.com/evrintobing17/MyGram/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/MyGram/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/MyGram/app/helpers/routehelper"
	"github.com/evrintobing17/MyGram/app/helpers/structsconverter"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	"github.com/evrintobing17/MyGram/app/modules/users"
	userDTO "github.com/evrintobing17/MyGram/app/modules/users/delivery/userDto"
	userUsecase "github.com/evrintobing17/MyGram/app/modules/users/usecase"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUC         users.UserUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewAuthHTTPHandler(r *gin.Engine, userUC users.UserUsecase, authMiddleware authmiddleware.AuthMiddleware) {
	handlers := userHandler{
		userUC:         userUC,
		authMiddleware: authMiddleware,
	}

	authorized := r.Group("/users")
	{
		authorized.POST("/login", handlers.login)
		authorized.POST("/register", handlers.register)
	}
	users := r.Group("/users", handlers.authMiddleware.AuthorizeJWTWithUserContext())
	{
		users.DELETE("", handlers.deleteUser)
		users.PUT("", handlers.updateUser)
	}
}

func (handler *userHandler) login(c *gin.Context) {

	var loginReq userDTO.ReqLogin

	errBind := c.ShouldBind(&loginReq)
	if errBind != nil {

		validations := requestvalidationerror.GetvalidationError(errBind)

		if len(validations) > 0 {
			jsonhttpresponse.BadRequest(c, validations)
			return
		}
		jsonhttpresponse.BadRequest(c, "")
		return
	}

	_, jwt, err := handler.userUC.Login(loginReq.Email, loginReq.Password)
	if err != nil {

		if err == userUsecase.ErrInvalidCredential {
			jsonhttpresponse.Unauthorized(c, jsonhttpresponse.NewFailedResponse(err.Error()))
			return
		}

		jsonhttpresponse.InternalServerError(c, jsonhttpresponse.NewFailedResponse(err.Error()))
		return
	}

	response := userDTO.ResLogin{
		Jwt: jwt,
	}
	jsonhttpresponse.OK(c, response)
	return
}

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

	user, err := handler.userUC.Register(registerReq.Username, registerReq.Email, registerReq.Password, registerReq.Age)

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

	response := userDTO.ResRegister{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
	}

	jsonhttpresponse.StatusCreated(c, response)
	return
}

func (handler *userHandler) updateUser(c *gin.Context) {
	var request userDTO.ReqUpdate
	errBind := c.ShouldBindJSON(&request)
	if errBind != nil {
		validations := requestvalidationerror.GetvalidationError(errBind)

		if len(validations) > 0 {
			jsonhttpresponse.BadRequest(c, validations)
			return
		}

		jsonhttpresponse.BadRequest(c, errBind.Error())
		return
	}

	updatedDriverData, err := structsconverter.ToMap(request)
	if err != nil {
		jsonhttpresponse.InternalServerError(c, err.Error())
	}

	//get user ID
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	updatedDriverData["id"] = userAuth.ID

	updatedUser, err := handler.userUC.UpdateUser(updatedDriverData)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	resp := userDTO.RespUpdate{
		ID:        updatedUser.ID,
		Email:     updatedUser.Email,
		Username:  updatedUser.Username,
		Age:       updatedUser.Age,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	jsonhttpresponse.OK(c, resp)
	return
}

func (handler *userHandler) deleteUser(c *gin.Context) {
	//get user ID
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	err = handler.userUC.DeleteUserByID(userAuth.ID)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
	}

	resp := userDTO.DeleteResp{
		Message: "Your account has been succesfully deleted",
	}

	jsonhttpresponse.OK(c, resp)
	return
}
