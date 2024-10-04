package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/package/elastic"
	"github.com/hritesh04/news-system/package/prometheus"
	"gorm.io/gorm"
)

type RestHandler struct {
	Router           *gin.Engine
	DB               *gorm.DB
	AuthService      *auth.Auth
	ElasticClient    *elastic.ElasticClient
	PrometheusClient *prometheus.PrometheusClient
}
