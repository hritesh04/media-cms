package repositories

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	result := r.db.Preload("Articles").Preload("Articles.Category").First(&user, "email = ?", email)
	if err := result.Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	result := r.db.Create(&user)
	if err := result.Error; err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

// func (r *cmsRepository) GetUserByID(id uint) (*domain.User, error) {
// 	user := new(domain.User)
// 	result := r.db.First(&user, "id = ?", id)
// 	if err := result.Error; err != nil {
// 		return user, nil
// 	}
// 	return user, nil
// }

func (r *userRepository) GetArticleByID(id string) (*domain.Article, error) {
	article := new(domain.Article)
	if err := r.db.First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}
