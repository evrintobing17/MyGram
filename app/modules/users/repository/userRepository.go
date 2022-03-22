package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB)
