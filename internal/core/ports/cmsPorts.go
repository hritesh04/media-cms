package ports

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	elastic "github.com/olivere/elastic/v7"
)

type CmsRepository interface {
	// GetUserByID(uint) (*domain.User, error)
	GetUserByEmail(string) (*domain.User, error)
	InsertUser(*domain.User) (*domain.User, error)
	GetArticleByID(string) (*domain.Article, error)
	UpdateArticle(*domain.Article) (*domain.Article, error)
	InsertArticle(*domain.Article) (*domain.Article, error)
	RemoveArticle(string) error
}

type CmsService interface {
	SignInUser(dto.LogInRequest) (*domain.User, error)
	CreateUser(dto.SignUpRequest) (*domain.User, error)
	GetArticleByID(string) (*domain.Article, error)
	UpdateArticle(dto.Article) (*domain.Article, error)
	CreateArticle(dto.Article) (*domain.Article, error)
	DeleteArticle(string) error
	SearchArticle(string) ([]*elastic.SearchHit, error)
}
