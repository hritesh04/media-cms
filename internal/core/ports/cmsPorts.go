package ports

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
)

type CmsRepository interface {
	InsertUser(*domain.User) (*domain.User, error)
	GetUserByEmail(string) (*domain.User, error)
}

type CmsService interface {
	CreateUser(dto.SignUpRequest) (*domain.User, error)
	SignInUser(dto.LogInRequest) (*domain.User, error)
}
