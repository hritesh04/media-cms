package services

import (
	"log"

	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	es "github.com/hritesh04/news-system/package/elastic"
	"github.com/lib/pq"
)

type articleService struct {
	articleRepository ports.ArticleRepository
	authService       ports.AuthService
	elasticClient     *es.ElasticClient
}

func NewArticleService(repository ports.ArticleRepository, authService ports.AuthService, elasticClient *es.ElasticClient) *articleService {
	return &articleService{
		articleRepository: repository,
		authService:       authService,
		elasticClient:     elasticClient,
	}
}

func (s *articleService) UpdateArticle(data dto.Article) (*domain.Article, error) {
	article := &domain.Article{
		Title:      data.Title,
		Content:    data.Content,
		Tags:       data.Tags,
		CategoryID: data.CategoryID,
		UserID:     data.UserId,
	}
	updatedArticle, err := s.articleRepository.UpdateArticle(article)
	if err != nil {
		return updatedArticle, err
	}
	return updatedArticle, nil
}

func (s *articleService) CreateArticle(data dto.Article) (*domain.Article, error) {
	article := &domain.Article{
		Title:      data.Title,
		Content:    data.Content,
		Tags:       pq.StringArray(data.Tags),
		CategoryID: data.CategoryID,
		UserID:     data.UserId,
	}
	newArticle, err := s.articleRepository.InsertArticle(article)
	if err != nil {
		return &domain.Article{}, nil
	}
	_, err = s.elasticClient.Index("article", newArticle)
	if err != nil {
		log.Println("Error indexing article in Elasticsearch:", err)
	}
	return newArticle, nil
}

func (s *articleService) DeleteArticle(id string) error {
	err := s.articleRepository.RemoveArticle(id)
	if err != nil {
		return err
	}
	_, err = s.elasticClient.Delete("article", id)
	if err != nil {
		log.Println("Error deleting article from Elasticsearch:", err)
	}
	return nil
}
