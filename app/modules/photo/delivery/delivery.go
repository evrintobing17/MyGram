package delivery

import (
	"strconv"

	"github.com/evrintobing17/MyGram/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/MyGram/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/MyGram/app/helpers/routehelper"
	"github.com/evrintobing17/MyGram/app/helpers/structsconverter"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	"github.com/evrintobing17/MyGram/app/modules/photo"
	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	PhotoUC        photo.PhotoUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewPhotoHTTPHandler(r *gin.Engine, photoUC photo.PhotoUsecase, authmiddleware authmiddleware.AuthMiddleware) {
	handlers := photoHandler{
		PhotoUC:        photoUC,
		authMiddleware: authmiddleware,
	}
	authorized := r.Group("/photos", handlers.authMiddleware.AuthorizeJWTWithUserContext())
	{
		authorized.POST("", handlers.addPhoto)
		authorized.GET("", handlers.getPhoto)
		authorized.PUT("/:photoId", handlers.updatePhoto)
		authorized.DELETE("/:photoId", handlers.deletePhoto)
	}
}

func (handler *photoHandler) addPhoto(c *gin.Context) {
	var request InsertReq

	errBind := c.ShouldBind(&request)
	if errBind != nil {

		validations := requestvalidationerror.GetvalidationError(errBind)

		if len(validations) > 0 {
			jsonhttpresponse.BadRequest(c, validations)
			return
		}
		jsonhttpresponse.BadRequest(c, "")
		return
	}

	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	insertPhoto, err := handler.PhotoUC.AddPhoto(request.Title, request.Caption, request.PhotoUrl, userAuth.ID)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	response := InsertResp{
		ID:        insertPhoto.ID,
		Title:     insertPhoto.Title,
		Caption:   insertPhoto.Caption,
		PhotoUrl:  insertPhoto.PhotoUrl,
		UserID:    insertPhoto.UserID,
		CreatedAt: insertPhoto.CreatedAt,
	}

	jsonhttpresponse.StatusCreated(c, response)

}

func (handler *photoHandler) getPhoto(c *gin.Context) {
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	getPhoto, err := handler.PhotoUC.GetPhoto(userAuth.ID, userAuth.Username, userAuth.Email)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}
	jsonhttpresponse.OK(c, getPhoto)
}

func (handler *photoHandler) updatePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	id, _ := strconv.Atoi(photoId)

	var request UpdateRequest
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

	updatedData, err := structsconverter.ToMap(request)
	if err != nil {
		jsonhttpresponse.InternalServerError(c, err.Error())
	}

	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	updatedData["user_id"] = userAuth.ID
	updatedData["id"] = id
	updatedPhoto, err := handler.PhotoUC.UpdatePhoto(updatedData)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err.Error())
		return
	}

	resp := UpdateResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoUrl:  updatedPhoto.PhotoUrl,
		UserID:    updatedPhoto.UserID,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	jsonhttpresponse.OK(c, resp)
}

func (handler *photoHandler) deletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	id, _ := strconv.Atoi(photoId)
	//get user ID
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	errs := handler.PhotoUC.DeleteUserByID(id, userAuth.ID)
	if errs != nil {
		jsonhttpresponse.BadRequest(c, errs.Error())
		return
	}

	resp := DeleteResp{
		Message: "Your account has been succesfully deleted",
	}

	jsonhttpresponse.OK(c, resp)
	return
}
