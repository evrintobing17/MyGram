package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/comment"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewCommentsRepository(db *gorm.DB) comment.CommentsRepository {
	return &repo{
		db: db,
	}
}

//Add new comments to db
func (r *repo) Insert(comments *models.Comments) (*models.Comments, error) {
	db := r.db.Create(&comments)
	if db.Error != nil {
		return nil, db.Error
	}
	return comments, nil
}

//Delete existing comments
func (r *repo) Delete(commentsId int) error {
	comments := models.Comments{ID: commentsId}
	db := r.db.Delete(&comments, "id = ?", commentsId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Get comments data by id
func (r *repo) GetAllById(userId int) (*[]models.Comments, error) {
	var comments []models.Comments

	db := r.db.Find(&comments, "user_id = ?", userId)
	if db.Error != nil {
		return nil, db.Error
	}
	return &comments, nil
}

func (r *repo) UpdatePartial(updateData map[string]interface{}) (*models.Comments, error) {
	id := updateData["id"]
	if id == nil {
		return nil, errors.New("field if cannot be empty")
	}
	idString := fmt.Sprintf("%v", id)
	driverID, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	var existingComments models.Comments
	db := r.db.First(&existingComments, "id=?", driverID)
	if db.Error != nil {
		return nil, db.Error
	}

	db = r.db.Debug().Model(&existingComments).Updates(updateData)
	if db.Error != nil {
		return nil, db.Error
	}

	return &existingComments, nil
}

func (r *repo) CheckIfUserIDExists(commentId, userID int) error {
	var comment models.Comments
	db := r.db.Find(&comment, "id = ? and user_id =?", commentId, userID)
	if db.Error != nil {
		return gorm.ErrRecordNotFound
	}
	return nil
}
