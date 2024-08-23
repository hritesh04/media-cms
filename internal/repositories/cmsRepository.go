package repositories

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"gorm.io/gorm"
)

type cmsRepository struct {
	DB *gorm.DB
}

func NewCms(db *gorm.DB) *cmsRepository {

	return &cmsRepository{
		DB: db,
	}
}

func (r *cmsRepository) InsertUser(user *domain.User) (*domain.User, error) {
	result := r.DB.Create(&user)
	if err := result.Error; err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
