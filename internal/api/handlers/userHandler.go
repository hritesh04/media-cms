package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/internal/api/rest"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	"github.com/hritesh04/news-system/internal/core/services"
	"github.com/hritesh04/news-system/internal/helper"
	"github.com/hritesh04/news-system/internal/repositories"
)

type handler struct {
	userService ports.UserService
}

func SetupUserRoutes(rh rest.RestHandler) {

	userRepo := repositories.NewUserRepository(rh.DB)
	svc := services.NewUserService(userRepo, rh.AuthService, rh.ElasticClient, rh.PrometheusClient)

	handler := &handler{
		userService: svc,
	}

	userGroup := rh.Router
	userGroup.POST("/signup", handler.SignUp)
	userGroup.POST("/login", handler.Login)
	userGroup.GET("/search", handler.SearchArticle)
	userGroup.GET("/:articleId", handler.GetArticleByID)
	// userGroup.GET("/profile", rh.GetProfile)
}

func (h *handler) SignUp(g *gin.Context) {
	var user dto.SignUpRequest
	if err := g.ShouldBindJSON(&user); err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err.Error())
	}
	token, err := h.userService.CreateUser(user)
	if err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err.Error())
	}
	g.SetCookie("media", token, 3600*24, "/", "localhost", false, true)
	helper.ReturnSuccess(g, http.StatusOK, token)
}

func (h *handler) Login(g *gin.Context) {
	var credentials dto.LogInRequest
	if err := g.ShouldBindJSON(&credentials); err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	token, err := h.userService.SignInUser(credentials)
	if err != nil {
		helper.ReturnFailed(g, http.StatusUnauthorized, err)
	}
	g.SetCookie("media", token, 3600*24, "/", "localhost", false, true)
	helper.ReturnSuccess(g, http.StatusOK, token)
}

func (h *handler) GetArticleByID(g *gin.Context) {
	articleID := g.Param("articleId")
	article, err := h.userService.GetArticleByID(articleID)
	if err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, article)
}

func (h *handler) SearchArticle(g *gin.Context) {
	query := g.Query("search")
	searchResult, err := h.userService.SearchArticle(query)
	if err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, searchResult)
}
