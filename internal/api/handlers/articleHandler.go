package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/internal/api/rest"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	"github.com/hritesh04/news-system/internal/core/services"
	"github.com/hritesh04/news-system/internal/helper"
	"github.com/hritesh04/news-system/internal/repositories"
)

type articleHandler struct {
	logger         *slog.Logger
	articleService ports.ArticleService
}

func SetupArticleRoutes(logger *slog.Logger, rh rest.RestHandler) {
	articleRepo := repositories.NewArticleRepository(rh.DB)
	svc := services.NewArticleService(articleRepo, rh.ElasticClient)
	handler := &articleHandler{
		logger:         logger,
		articleService: svc,
	}
	articleGroup := rh.Router.Group("/article")
	articleGroup.Use(rh.AuthService.IsAuthor())
	articleGroup.POST("/", handler.CreateArticle)
	articleGroup.PATCH("/:articleId", handler.UpdateArticle)
	articleGroup.DELETE("/:articleId", handler.DeleteArticle)
}

func (h *articleHandler) CreateArticle(g *gin.Context) {
	var article dto.Article
	if err := g.ShouldBindJSON(&article); err != nil {
		h.logger.Error("request body parsing", "error", err)
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	userString := g.GetHeader("userID")
	userId, err := strconv.ParseUint(userString, 10, 32)
	if err != nil {
		h.logger.Error("fetching userID from header", "error", err)
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	article.UserId = uint(userId)
	result, err := h.articleService.CreateArticle(article)
	if err != nil {
		h.logger.Error("creating article", "error", err)
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *articleHandler) UpdateArticle(g *gin.Context) {
	var article dto.Article
	if err := g.ShouldBindJSON(&article); err != nil {
		h.logger.Error("request body parsing", "error", err)
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	articleId, err := strconv.ParseUint(g.Param("articleId"), 10, 32)
	if err != nil {
		h.logger.Error("parsing articleID from param", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	article.ID = uint(articleId)
	result, err := h.articleService.UpdateArticle(article)
	if err != nil {
		h.logger.Error("updating article", "error", err)
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *articleHandler) DeleteArticle(g *gin.Context) {
	if err := h.articleService.DeleteArticle(g.Param("articleId")); err != nil {
		h.logger.Error("deleting article by ID", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, "Article deleted sucessfully")
}
