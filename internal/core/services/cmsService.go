package services

import "github.com/hritesh04/news-system/internal/core/ports"

type cmsService struct {
	cmsRepository ports.CmsRepository
}

func NewCmsService(repository ports.CmsRepository) *cmsService {
	return &cmsService{
		cmsRepository: repository,
	}
}

func (s *cmsService) AddUser() {

}
