package domain

import (
	"database/sql/driver"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role string

const (
	AUTHOR Role = "AUTHOR"
	USER   Role = "USER"
)

func (r *Role) Scan(value string) error {
	*r = Role(value)
	return nil
}

func (r Role) Value() driver.Value {
	return string(r)
}

// TODO: add default values
type User struct {
	gorm.Model
	Name          string
	Email         string
	Password      string
	Type          Role           `gorm:"type:role"`
	Articles      []Article      `gorm:"foreignKey:UserID"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID"`
}

type Article struct {
	gorm.Model
	Title      string
	Content    string
	Tags       pq.StringArray `gorm:"type:text[]"`
	Comments   []Comment      `gorm:"foreignKey:ArticleID"`
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	Content   string
	ArticleID uint
	On        Article `gorm:"foreignKey:ArticleID"`
	UserID    uint
	By        User `gorm:"foreignKey:UserID"`
}

type Subscription struct {
	gorm.Model
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
}

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Articles []Article `gorm:"foreignKey:CategoryID"`
}
