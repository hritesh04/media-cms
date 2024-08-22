package repositories

import (
	"gorm.io/gorm"
)

type cmsRepository struct {
	DB *gorm.DB
}

func NewCms(db *gorm.DB) *cmsRepository {

	return &cmsRepository{
		DB: db,
	}
}

func (r *cmsRepository) AddUser() {

}
