package domain

import (
	"database/sql/driver"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role string

const (
	AUTHOR Role = "author"
	USER   Role = "user"
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
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	Type          Role           `json:"type" gorm:"type:role;default:user"`
	Articles      []Article      `json:"articles" gorm:"foreignKey:UserID"`
	Subscriptions []Subscription `json:"subscriptions" gorm:"foreignKey:UserID"`
}

type Article struct {
	gorm.Model
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Tags       pq.StringArray `json:"tags" gorm:"type:text[]"`
	Comments   []Comment      `json:"comments" gorm:"foreignKey:ArticleID"`
	CategoryID uint           `json:"category_id"`
	Category   Category       `json:"-" gorm:"foreignKey:CategoryID"`
	UserID     uint           `json:"user_id"`
	User       User           `json:"-" gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	Content   string  `json:"content"`
	ArticleID uint    `json:"article_id"`
	On        Article `gorm:"foreignKey:ArticleID"`
	UserID    uint    `json:"user_id"`
	By        User    `gorm:"foreignKey:UserID"`
}

type Subscription struct {
	gorm.Model `json:"-"`
	AuthorID   uint     `json:"author_id"`
	Author     User     `json:"author" gorm:"foreignKey:AuthorID"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	UserID     uint     `json:"user_id"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
}

type Category struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Articles []Article `json:"articles" gorm:"foreignKey:CategoryID"`
}
