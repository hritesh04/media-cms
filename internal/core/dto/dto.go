package dto

// General
// Culture
// Arts
// Geography
// HealthAndFitness
// History
// Psychology
// Mathematics
// Natural
// Lifestyle
// Philosophy
// SocialScience
// Technology

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Article struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Tags       []string `json:"tags"`
	CategoryID uint     `json:"category"`
	UserId     uint
}

type Comment struct {
	Content   string `json:"content"`
	ArticleID uint   `json:"article_id"`
	UserID    uint   `json:"user_id"`
}

type SubscriptionRequest struct {
	AuthorID   uint `json:"author_id"`
	CategoryID uint `json:"category_id"`
	UserID     uint `json:"user_id"`
}
