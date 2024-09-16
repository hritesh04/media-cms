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
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Article struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Tags       []string `json:"tags"`
	CategoryID uint     `json:"category"`
	UserId     uint
}
