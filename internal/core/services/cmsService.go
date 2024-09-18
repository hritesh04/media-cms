package services

import (
	"fmt"

	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	"github.com/lib/pq"
)

type cmsService struct {
	cmsRepository ports.CmsRepository
}

func NewCmsService(repository ports.CmsRepository) *cmsService {
	return &cmsService{
		cmsRepository: repository,
	}
}

func (s *cmsService) SignInUser(data dto.LogInRequest) (*domain.User, error) {
	user, err := s.cmsRepository.GetUserByEmail(data.Email)
	if err != nil {
		return user, err
	}
	if success := auth.ComparePassword(user.Password, data.Password); !success {
		return &domain.User{}, fmt.Errorf("incorrect password")
	}
	return user, nil
}

func (s *cmsService) CreateUser(data dto.SignUpRequest) (*domain.User, error) {
	hash, err := auth.HashPassword(data.Password)
	if err != nil {
		return &domain.User{}, err
	}
	user := &domain.User{
		Name:     data.Name,
		Password: hash,
		Type:     domain.AUTHOR,
		Email:    data.Email,
	}
	newUser, err := s.cmsRepository.InsertUser(user)
	if err != nil {
		return &domain.User{}, err
	}
	return newUser, nil
}

func (s *cmsService) GetArticleByID(id string) (*domain.Article, error) {
	article, err := s.cmsRepository.GetArticleByID(id)

	if err != nil {
		return article, err
	}

	return article, nil
}

func (s *cmsService) UpdateArticle(data dto.Article) (*domain.Article, error) {
	article := &domain.Article{
		Title:      data.Title,
		Content:    data.Content,
		Tags:       data.Tags,
		CategoryID: data.CategoryID,
		UserID:     data.UserId,
	}
	updatedArticle, err := s.cmsRepository.UpdateArticle(article)
	if err != nil {
		return updatedArticle, err
	}
	return updatedArticle, nil
}

func (s *cmsService) CreateArticle(data dto.Article) (*domain.Article, error) {
	article := &domain.Article{
		Title:      data.Title,
		Content:    data.Content,
		Tags:       pq.StringArray(data.Tags),
		CategoryID: data.CategoryID,
		UserID:     data.UserId,
	}
	newArticle, err := s.cmsRepository.InsertArticle(article)
	if err != nil {
		return &domain.Article{}, nil
	}
	return newArticle, nil
}

func (s *cmsService) DeleteArticle(string) error {
	return nil
}
