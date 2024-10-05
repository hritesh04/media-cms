package services

import (
	"fmt"
	"strconv"

	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	es "github.com/hritesh04/news-system/package/elastic"
	"github.com/hritesh04/news-system/package/prometheus"
	"github.com/olivere/elastic/v7"
)

type userService struct {
	userRepository   ports.UserRepository
	Auth             ports.AuthService
	elasticClient    *es.ElasticClient
	prometheusClient *prometheus.PrometheusClient
}

func NewUserService(repository ports.UserRepository, authService ports.AuthService, elasticClient *es.ElasticClient, prometheusClient *prometheus.PrometheusClient) *userService {
	return &userService{
		userRepository:   repository,
		Auth:             authService,
		elasticClient:    elasticClient,
		prometheusClient: prometheusClient,
	}
}

func (s *userService) SignInUser(data dto.LogInRequest) (string, error) {
	user, err := s.userRepository.GetUserByEmail(data.Email)
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

func (s *userService) SignUpUser(data dto.SignUpRequest) (string, error) {
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
	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	token, err := s.Auth.GenerateToken(newUser.ID, newUser.Type)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) GetArticleByID(id string) (*domain.Article, error) {
	article, err := s.userRepository.GetArticleByID(id)
	s.prometheusClient.Increment("article_views", strconv.Itoa(int(article.ID)))
	if err != nil {
		return article, err
	}

	return article, nil
}

func (s *userService) SearchArticle(query string) ([]*elastic.SearchHit, error) {
	searchResult, err := s.elasticClient.Search("article", elastic.NewMatchQuery("Title", query))
	if err != nil {
		return nil, err
	}
	return searchResult.Hits.Hits, nil
}
