package ports

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
)

type CmsRepository interface {
	GetUserByID(uint) (*domain.User, error)
	GetUserByEmail(string) (*domain.User, error)
	InsertUser(*domain.User) (*domain.User, error)
	InsertArticle(*domain.Article) (*domain.Article, error)
	GetArticleByID(string) (*domain.Article, error)
}

type CmsService interface {
	CreateUser(dto.SignUpRequest) (*domain.User, error)
	SignInUser(dto.LogInRequest) (*domain.User, error)
	CreateArticle(dto.Article) (*domain.Article, error)
	GetArticleByID(string) (*domain.Article, error)
}
