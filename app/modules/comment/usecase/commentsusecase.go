package usecase

import (
	"time"

	"github.com/evrintobing17/MyGram/app/models"
	"github.com/evrintobing17/MyGram/app/modules/comment"
	"github.com/evrintobing17/MyGram/app/modules/photo"
)

type GetCommentResp struct {
	ID        int        `json:"id"`
	Message   string     `json:"message"`
	PhotoID   int        `json:"photo_id"`
	UserID    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	User      `json:"User"`
	Photo     `json:"Photo"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Photo struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

type UC struct {
	repo      comment.CommentsRepository
	photoRepo photo.PhotosRepository
}

func NewCommentUsecase(repo comment.CommentsRepository, photoRepo photo.PhotosRepository) comment.CommentUsecase {
	return &UC{
		repo:      repo,
		photoRepo: photoRepo,
	}
}

func (uc *UC) AddComment(message string, photoId, userId int) (*models.Comments, error) {
	modelsComment := models.Comments{
		UserID:   userId,
		PhotosID: photoId,
		Message:  message,
	}

	data, err := uc.repo.Insert(&modelsComment)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *UC) GetComment(userID int, username, email string) (interface{}, error) {
	var resp []GetCommentResp

	data, err := uc.repo.GetAllById(userID)
	if err != nil {
		return nil, err
	}

	for _, comment := range *data {
		dataPhoto, err := uc.photoRepo.GetByUserID(comment.PhotosID)
		if err != nil {
			return nil, err
		}
		resp = append(resp, GetCommentResp{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotosID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: User{
				ID:       userID,
				Username: username,
				Email:    email,
			},
			Photo: Photo{
				ID:        dataPhoto.ID,
				Title:     dataPhoto.Title,
				Caption:   dataPhoto.Caption,
				PhotoUrl:  dataPhoto.PhotoUrl,
				UserID:    dataPhoto.UserID,
				UpdatedAt: dataPhoto.UpdatedAt,
				CreatedAt: dataPhoto.CreatedAt,
			},
		})

	}
	return resp, nil
}

func (uc *UC) DeleteCommentByID(commentId, userId int) error {

	err := uc.repo.CheckIfUserIDExists(commentId, userId)
	if err != nil {
		return err
	}

	errDelete := uc.repo.Delete(commentId)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func (uc *UC) UpdateComment(updateData map[string]interface{}) (*models.Comments, error) {

	userId := updateData["user_id"].(int)
	photoId := updateData["id"].(int)

	err := uc.repo.CheckIfUserIDExists(photoId, userId)
	if err != nil {
		return nil, err
	}

	updateData["updated_at"] = time.Now()
	userData, err := uc.repo.UpdatePartial(updateData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
