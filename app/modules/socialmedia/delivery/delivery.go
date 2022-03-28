package delivery

import (
	"strconv"

	"github.com/evrintobing17/MyGram/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/MyGram/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/MyGram/app/helpers/routehelper"
	"github.com/evrintobing17/MyGram/app/helpers/structsconverter"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	"github.com/evrintobing17/MyGram/app/modules/socialmedia"
	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaUC  socialmedia.SocialMediaUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewSocialMediaHandler(r *gin.Engine, socialMediaUC socialmedia.SocialMediaUsecase,
	authMiddleware authmiddleware.AuthMiddleware) {
	handler := socialMediaHandler{
		socialMediaUC:  socialMediaUC,
		authMiddleware: authMiddleware,
	}
	socialMedia := r.Group("/socialmedias", handler.authMiddleware.AuthorizeJWTWithUserContext())
	{
		socialMedia.GET("", handler.getSocialMedia)
		socialMedia.POST("", handler.addSocialMedia)
		socialMedia.DELETE("/:socialMediaId", handler.deleteComment)
		socialMedia.PUT("/:socialMediaId", handler.updateComment)
	}
}

func (handler *socialMediaHandler) getSocialMedia(c *gin.Context) {
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	getcomment, err := handler.socialMediaUC.GetSocialMedia(userAuth.ID, userAuth.Username)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}
	jsonhttpresponse.OK(c, getcomment)
}

func (handler *socialMediaHandler) addSocialMedia(c *gin.Context) {
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

	insertSocialMedia, err := handler.socialMediaUC.AddSocialMedia(request.Name, request.SocialMediaUrl, userAuth.ID)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	response := InsertResp{
		ID:             insertSocialMedia.ID,
		Name:           insertSocialMedia.Name,
		SocialMediaUrl: insertSocialMedia.SocialMediaUrl,
		UserID:         insertSocialMedia.UserID,
		CreatedAt:      insertSocialMedia.CreatedAt,
	}

	jsonhttpresponse.StatusCreated(c, response)
}

func (handler *socialMediaHandler) updateComment(c *gin.Context) {
	commentId := c.Param("socialMediaId")

	id, _ := strconv.Atoi(commentId)

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
	updatedComment, err := handler.socialMediaUC.UpdateSocialMedia(updatedData)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err.Error())
		return
	}

	resp := UpdateResponse{
		ID:             updatedComment.ID,
		Name:           updatedComment.Name,
		SocialMediaUrl: updatedComment.SocialMediaUrl,
		UserID:         updatedComment.UserID,
		UpdateedAt:     updatedComment.UpdatedAt,
	}

	jsonhttpresponse.OK(c, resp)
}

func (handler *socialMediaHandler) deleteComment(c *gin.Context) {
	commentId := c.Param("socialMediaId")

	id, _ := strconv.Atoi(commentId)
	//get user ID
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	errs := handler.socialMediaUC.DeleteSocialMediaByID(id, userAuth.ID)
	if errs != nil {
		jsonhttpresponse.BadRequest(c, errs.Error())
		return
	}

	resp := DeleteResp{
		Message: "Your account has been succesfully deleted",
	}

	jsonhttpresponse.OK(c, resp)
}
