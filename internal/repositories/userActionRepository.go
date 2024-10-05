package repositories

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"gorm.io/gorm"
)

type userActionRepository struct {
	db *gorm.DB
}

func NewUserActionRepository(db *gorm.DB) *userActionRepository {
	return &userActionRepository{
		db: db,
	}
}

func (r *userActionRepository) CreateComment(data *domain.Comment) (*domain.Comment, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return &domain.Comment{}, err
	}
	return data, nil
}

func (r *userActionRepository) DeleteComment(commentID string) error {
	if err := r.db.Delete(&domain.Comment{}, "id = ?", commentID).Error; err != nil {
		return err
	}
	return nil
}

func (r *userActionRepository) AddSubscription(data *domain.Subscription) (*domain.Subscription, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return &domain.Subscription{}, nil
	}
	return data, nil
}

func (r *userActionRepository) RemoveSubscription(subscriptionID string) error {
	if err := r.db.Delete(&domain.Subscription{}, "id = ?", subscriptionID).Error; err != nil {
		return err
	}
	return nil
}
