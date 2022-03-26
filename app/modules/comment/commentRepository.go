package comment

import "github.com/evrintobing17/MyGram/app/models"

type CommentsRepository interface {
	Insert(comment *models.Comments) (*models.Comments, error)
	Delete(commentId int) error
	GetAllById(user int) (*[]models.Comments, error)
	UpdatePartial(updateData map[string]interface{}) (*models.Comments, error)
	CheckIfUserIDExists(commentId, userID int) error
}
