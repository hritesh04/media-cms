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

func (s *cmsService) CreateArticle(data dto.Article) (*domain.Article, error) {
	// user, err := s.cmsRepository.GetUserByID(data.UserId)
	// if err != nil {
	// 	return &domain.Article{}, err
	// }
	// if user.Type != "AUTHOR" {
	// 	return &domain.Article{}, fmt.Errorf("user is not an author")
	// }
	article := &domain.Article{
		Title:      data.Title,
		Content:    data.Content,
		Tags:       pq.StringArray(data.Tags),
		CategoryID: data.CategoryID,
		UserID:     data.UserId,
	}
	fmt.Println(article)
	newArticle, err := s.cmsRepository.InsertArticle(article)
	if err != nil {
		return &domain.Article{}, nil
	}
	return newArticle, nil
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

func (s *cmsService) GetArticleByID(id string) (*domain.Article, error) {
	article, err := s.cmsRepository.GetArticleByID(id)

	if err != nil {
		return article, err
	}

	return article, nil
}
