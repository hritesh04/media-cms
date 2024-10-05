package services

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
)

type userActionService struct {
	repo ports.UserActionRepository
}

func NewUserActionService(repo ports.UserActionRepository) *userActionService {
	return &userActionService{
		repo: repo,
	}
}

func (s *userActionService) AddComment(data dto.Comment) (*domain.Comment, error) {
	comment := &domain.Comment{
		Content:   data.Content,
		ArticleID: data.ArticleID,
		UserID:    data.UserID,
	}
	result, err := s.repo.CreateComment(comment)
	if err != nil {
		return &domain.Comment{}, err
	}
	return result, nil
}

func (s *userActionService) RemoveComment(commentID string) error {
	if err := s.repo.DeleteComment(commentID); err != nil {
		return err
	}
	return nil
}

func (s *userActionService) Subscribe(data dto.SubscriptionRequest) (*domain.Subscription, error) {
	subscription := &domain.Subscription{
		AuthorID:   data.AuthorID,
		CategoryID: data.CategoryID,
		UserID:     data.UserID,
	}
	result, err := s.repo.AddSubscription(subscription)
	if err != nil {
		return &domain.Subscription{}, nil
	}
	return result, nil
}

func (s *userActionService) UnSubscribe(subscriptionID string) error {
	if err := s.repo.RemoveSubscription(subscriptionID); err != nil {
		return err
	}
	return nil
}
