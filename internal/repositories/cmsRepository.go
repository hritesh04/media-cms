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

func (r *cmsRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	result := r.db.Preload("Articles").Preload("Articles.Category").First(&user, "email = ?", email)
	if err := result.Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *cmsRepository) InsertUser(user *domain.User) (*domain.User, error) {
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

func (r *cmsRepository) GetArticleByID(id string) (*domain.Article, error) {
	article := new(domain.Article)
	if err := r.db.First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *cmsRepository) UpdateArticle(data *domain.Article) (*domain.Article, error) {
	article := new(domain.Article)
	if err := r.db.First(&article).Error; err != nil {
		return article, err
	}
	article.Content = data.Content
	article.Title = data.Title
	article.Tags = data.Tags
	if err := r.db.Save(article).Error; err != nil {
		return article, err
	}
	return article, nil
}

func (r *cmsRepository) InsertArticle(article *domain.Article) (*domain.Article, error) {
	if err := r.db.Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *cmsRepository) RemoveArticle(id string) error {
	if err := r.db.Delete(&domain.Article{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
