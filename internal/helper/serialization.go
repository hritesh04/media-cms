package helper

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
)

func UserDtoToDomain(data *dto.SignUpRequest) *domain.User {
	return &domain.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
