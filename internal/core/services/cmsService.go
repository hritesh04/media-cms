package services

import (
	"fmt"
	"log"

	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	es "github.com/hritesh04/news-system/package/elastic"
	"github.com/lib/pq"
	"github.com/olivere/elastic/v7"
)

type cmsService struct {
	cmsRepository ports.CmsRepository
	Auth          ports.AuthService
	elasticClient *es.ElasticClient
}

func NewCmsService(repository ports.CmsRepository, authService ports.AuthService, elasticClient *es.ElasticClient) *cmsService {
	return &cmsService{
		cmsRepository: repository,
		Auth:          authService,
		elasticClient: elasticClient,
	}
}

func (s *cmsService) SignInUser(data dto.LogInRequest) (string, error) {
	user, err := s.cmsRepository.GetUserByEmail(data.Email)
	if err != nil {
		return "", err
	}
	if success := s.Auth.ComparePassword(user.Password, data.Password); !success {
		return "", fmt.Errorf("incorrect password")
	}
	token, err := s.Auth.GenerateToken(user.ID, user.Type)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *cmsService) CreateUser(data dto.SignUpRequest) (string, error) {
	hash, err := s.Auth.HashPassword(data.Password)
	if err != nil {
		return "", err
	}
	user := &domain.User{
		Name:     data.Name,
		Password: hash,
		Type:     domain.AUTHOR,
		Email:    data.Email,
	}
	newUser, err := s.cmsRepository.InsertUser(user)
	if err != nil {
		return "", err
	}
	token, err := s.Auth.GenerateToken(newUser.ID, newUser.Type)
	if err != nil {
		return "", err
	}
	return token, nil
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
	_, err = s.elasticClient.Index("article", newArticle)
	if err != nil {
		log.Println("Error indexing article in Elasticsearch:", err)
	}
	return newArticle, nil
}

func (s *cmsService) DeleteArticle(id string) error {
	err := s.cmsRepository.RemoveArticle(id)
	if err != nil {
		return err
	}
	_, err = s.elasticClient.Delete("article", id)
	if err != nil {
		log.Println("Error deleting article from Elasticsearch:", err)
	}
	return nil
}

func (s *cmsService) SearchArticle(query string) ([]*elastic.SearchHit, error) {
	searchResult, err := s.elasticClient.Search("article", elastic.NewMatchQuery("Title", query))
	if err != nil {
		return nil, err
	}
	return searchResult.Hits.Hits, nil
}
