package comment

import "github.com/evrintobing17/MyGram/app/models"

type CommentUsecase interface {
	AddComment(message string, photoId, userId int) (*models.Comments, error)
	GetComment(userID int, username, email string) (interface{}, error)
	DeleteCommentByID(commentId, userId int) error
	UpdateComment(updateData map[string]interface{}) (*models.Comments, error)
}
