package services

import (
	"fmt"

	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
)

type cmsService struct {
	cmsRepository ports.CmsRepository
}

func NewCmsService(repository ports.CmsRepository) *cmsService {
	return &cmsService{
		cmsRepository: repository,
	}
}

func (s *cmsService) CreateUser(data dto.SignUpRequest) (*domain.User, error) {
	hash, err := auth.HashPassword(data.Password)
	if err != nil {
		return &domain.User{}, err
	}
	user := domain.User{
		Name:     data.Username,
		Password: hash,
		Email:    data.Email,
	}
	newUser, err := s.cmsRepository.InsertUser(&user)
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
