package delivery

import (
	"strconv"

	"github.com/evrintobing17/MyGram/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/MyGram/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/MyGram/app/helpers/routehelper"
	"github.com/evrintobing17/MyGram/app/helpers/structsconverter"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	"github.com/evrintobing17/MyGram/app/modules/comment"
	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	CommentUC      comment.CommentUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewcommentHTTPHandler(r *gin.Engine, commentUC comment.CommentUsecase, authmiddleware authmiddleware.AuthMiddleware) {
	handlers := commentHandler{
		CommentUC:      commentUC,
		authMiddleware: authmiddleware,
	}
	authorized := r.Group("/comment", handlers.authMiddleware.AuthorizeJWTWithUserContext())
	{
		authorized.POST("", handlers.addComment)
		authorized.GET("", handlers.getComment)
		authorized.PUT("/:commentId", handlers.updateComment)
		authorized.DELETE("/:commentId", handlers.deleteComment)
	}
}

// Add Comment godoc
// @Summary  Add Comment
// @ID addComment
// @Description add comment for mygram. Return comment model
// @Description add comment data
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param access-token header string true "Bearer Token"
// @Params authinfo body InsertReq true "add comment data"
// @Success 200 {object} InsertResp
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Security JWTToken
// @Router /comment [post]
func (handler *commentHandler) addComment(c *gin.Context) {
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

	insertComment, err := handler.CommentUC.AddComment(request.Message, request.PhotoID, userAuth.ID)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	response := InsertResp{
		ID:        insertComment.ID,
		Message:   insertComment.Message,
		PhotoID:   insertComment.PhotosID,
		UserID:    insertComment.UserID,
		CreatedAt: insertComment.CreatedAt,
	}

	jsonhttpresponse.StatusCreated(c, response)

}

// Get Comment godoc
// @Summary  Get Comment
// @ID getComment
// @Description get comment for mygram. Return comment model
// @Description get comment data
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param access-token header string true "Bearer Token"
// @Success 200 {object} interface{}
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Security JWTToken
// @Router /comment [get]
func (handler *commentHandler) getComment(c *gin.Context) {
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	getcomment, err := handler.CommentUC.GetComment(userAuth.ID, userAuth.Username, userAuth.Email)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}
	jsonhttpresponse.OK(c, getcomment)
}

// Update Comment godoc
// @Summary  Update Comment
// @ID updateComment
// @Description update data comment for mygram. Return comment model
// @Description update comment data
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param access-token header string true "Bearer Token"
// @Params authinfo body UpdateRequest true "add comment data"
// @Param id path integer true "commentId"
// @Success 200 {object} UpdateResponse
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Security JWTToken
// @Router /comment/{commentId} [put]
func (handler *commentHandler) updateComment(c *gin.Context) {
	commentId := c.Param("commentId")

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
	updatedComment, err := handler.CommentUC.UpdateComment(updatedData)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err.Error())
		return
	}

	resp := UpdateResponse{
		ID:        updatedComment.ID,
		Message:   updatedComment.Message,
		PhotoID:   updatedComment.PhotosID,
		UserID:    updatedComment.UserID,
		UpdatedAt: updatedComment.UpdatedAt,
	}

	jsonhttpresponse.OK(c, resp)
}

// Delete Comment godoc
// @Summary  Delete Comment
// @ID deleteComment
// @Description delete data comment for mygram. Return comment model
// @Description delete comment data
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param access-token header string true "Bearer Token"
// @Param id path integer true "commentId"
// @Success 200 {object} DeleteResp
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Security JWTToken
// @Router /comment/{commentId} [delete]
func (handler *commentHandler) deleteComment(c *gin.Context) {
	commentId := c.Param("commentId")

	id, _ := strconv.Atoi(commentId)
	//get user ID
	userAuth, err := routehelper.GetUserFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, err)
		return
	}

	errs := handler.CommentUC.DeleteCommentByID(id, userAuth.ID)
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
