package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	SignInUser(dto.LogInRequest) (string, error)
	CreateUser(dto.SignUpRequest) (string, error)
	GetArticleByID(string) (*domain.Article, error)
	UpdateArticle(dto.Article) (*domain.Article, error)
	CreateArticle(dto.Article) (*domain.Article, error)
	DeleteArticle(string) error
	SearchArticle(string) ([]*elastic.SearchHit, error)
}

type AuthService interface {
	IsAuthor() gin.HandlerFunc
	Authorize() gin.HandlerFunc
	ValidateUser(string) (jwt.MapClaims, error)
	GenerateToken(uint, domain.Role) (string, error)
	HashPassword(string) (string, error)
	ComparePassword(string, string) bool
}
