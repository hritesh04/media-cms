package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/internal/core/services"
	"github.com/hritesh04/news-system/internal/handlers"
	"github.com/hritesh04/news-system/internal/migrations"
	"github.com/hritesh04/news-system/internal/repositories"
	"github.com/hritesh04/news-system/package/elastic"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("HOST"), os.Getenv("USER_NAME"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed")
	}
	migrations.InitCmsMigrate(db)
	cmsRepository := repositories.NewCms(db)
	elasticClient := elastic.NewElasticClient(os.Getenv("ELASTICSEARCH_URL"))
	cmsService := services.NewCmsService(cmsRepository, elasticClient)
	cmsHandler := handlers.NewCmsHandler(cmsService)

	router := gin.New()
	router.Use(gin.Logger())
	router.RedirectTrailingSlash = false
	router.GET("/search", cmsHandler.SearchArticle)
	router.POST("/signup", cmsHandler.SignUp)
	router.POST("/login", cmsHandler.Login)
	router.GET("/read/:articleId", cmsHandler.GetArticle)
	router.Use(auth.Authorize())
	articleRouter := router.Group("/article")
	articleRouter.Use(auth.IsAuthor())
	articleRouter.POST("/", cmsHandler.CreateArticle)
	articleRouter.PUT("/:articleId", cmsHandler.UpdateArticle)
	router.Run(":3000")
}
