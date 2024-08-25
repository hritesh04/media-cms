package repositories

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"gorm.io/gorm"
)

type cmsRepository struct {
	db *gorm.DB
}

func NewCms(db *gorm.DB) *cmsRepository {

	return &cmsRepository{
		db: db,
	}
}

func (r *cmsRepository) InsertUser(user *domain.User) (*domain.User, error) {
	result := r.db.Create(&user)
	if err := result.Error; err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (r *cmsRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "email = ?", email)
	if err := result.Error; err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}
