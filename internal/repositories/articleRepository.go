package repositories

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) InsertArticle(article *domain.Article) (*domain.Article, error) {
	if err := r.db.Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *articleRepository) RemoveArticle(id string) error {
	if err := r.db.Delete(&domain.Article{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) UpdateArticle(data *domain.Article) (*domain.Article, error) {
	article := new(domain.Article)
	// if err := r.db.First(&article).Error; err != nil {
	// 	return article, err
	// }
	// article.Content = data.Content
	// article.Title = data.Title
	// article.Tags = data.Tags
	// if err := r.db.Save(article).Error; err != nil {
	// 	return article, err
	// }

	// better approach than above since it does less db calls and can return updated data in just one query.
	err := r.db.Model(article).Clauses(clause.Returning{}).Where("id=?", data.ID).Updates(data).Error
	if err != nil {
		return article, err
	}
	return article, nil
}
