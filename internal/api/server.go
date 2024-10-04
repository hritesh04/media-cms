package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/config"
	"github.com/hritesh04/news-system/internal/api/handlers"
	"github.com/hritesh04/news-system/internal/api/rest"
	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/internal/core/domain"
	"github.com/hritesh04/news-system/package/elastic"
	"github.com/hritesh04/news-system/package/prometheus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(cfg config.AppConfig) {

	router := gin.New()
	router.Use(gin.Logger())

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB connection failed: %v", err)
		panic("DB connection failed")
	}
	if err := db.AutoMigrate(&domain.User{}, &domain.Article{}, &domain.Category{}, &domain.Comment{}, &domain.Subscription{}); err != nil {
		log.Printf("DB migration failed: %v", err)
	}

	authService := auth.NewAuthService()
	elasticClient := elastic.NewElasticClient(cfg.ElasticUrl)
	prometheusClient := prometheus.NewPrometheusClient(cfg.PrometheusUrl)

	rh := rest.RestHandler{
		Router:           router,
		DB:               db,
		AuthService:      authService,
		ElasticClient:    elasticClient,
		PrometheusClient: prometheusClient,
	}

	setupRoutes(rh)

	router.Run(cfg.ServerPort)

}

func setupRoutes(rh rest.RestHandler) {
	rh.Router.GET("/metrics", rh.PrometheusClient.Handler())
	handlers.SetupUserRoutes(rh)
	rh.Router.Use(rh.AuthService.Authorize())
	handlers.SetupArticleRoutes(rh)
	// handlers.SetupCategoryRoutes(rh)
	// handlers.SetupCommentRoutes(rh)
	// handlers.SetupSubscriptionRoutes(rh)
}
