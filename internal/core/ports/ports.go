package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	elastic "github.com/olivere/elastic/v7"
)

type UserRepository interface {
	GetUserByEmail(string) (*domain.User, error)
	CreateUser(*domain.User) (*domain.User, error)
	GetAllArticle(int, int) (*[]domain.Article, error)
	GetArticleByID(string) (*domain.Article, error)
}

type UserService interface {
	SignInUser(dto.LogInRequest) (string, error)
	SignUpUser(dto.SignUpRequest) (string, error)
	GetAllArticle(string, string) (*[]domain.Article, error)
	GetArticleByID(string) (*domain.Article, error)
	SearchArticle(string) ([]*elastic.SearchHit, error)
}

type UserActionRepository interface {
	CreateComment(*domain.Comment) (*domain.Comment, error)
	DeleteComment(string) error
	AddSubscription(*domain.Subscription) (*domain.Subscription, error)
	RemoveSubscription(string) error
}

type UserActionService interface {
	AddComment(dto.Comment) (*domain.Comment, error)
	RemoveComment(string) error
	Subscribe(dto.SubscriptionRequest) (*domain.Subscription, error)
	UnSubscribe(string) error
}

type ArticleRepository interface {
	// GetArticleByID(string) (*domain.Article, error)
	InsertArticle(*domain.Article) (*domain.Article, error)
	UpdateArticle(*domain.Article) (*domain.Article, error)
	RemoveArticle(string) error
}

type ArticleService interface {
	UpdateArticle(dto.Article) (*domain.Article, error)
	CreateArticle(dto.Article) (*domain.Article, error)
	DeleteArticle(string) error
}

type AuthService interface {
	Authorize() gin.HandlerFunc
	IsAuthor() gin.HandlerFunc
	GenerateToken(uint, domain.Role) (string, error)
	ValidateUser(string) (jwt.MapClaims, error)
	HashPassword(string) (string, error)
	ComparePassword(string, string) bool
}
