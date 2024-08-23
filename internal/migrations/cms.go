package migrations

import (
	"github.com/hritesh04/news-system/internal/core/domain"
	"gorm.io/gorm"
)

func InitCmsMigrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{}, &domain.Article{}, &domain.Category{}, &domain.Comment{}, &domain.Subscription{})
}
