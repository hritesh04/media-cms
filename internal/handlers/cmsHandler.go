package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
)

type cmsHandler struct {
	cmsService ports.CmsService
}

func NewCmsHandler(service ports.CmsService) *cmsHandler {
	return &cmsHandler{
		cmsService: service,
	}
}

func (h *cmsHandler) SignUp(g *gin.Context) {
	var user dto.SignUpRequest
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"succcess": false,
			"data":     err.Error(),
		})
	}
	result, err := h.cmsService.CreateUser(user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"succcess": false,
			"data":     err.Error(),
		})
	}
	token, err := auth.GenerateToken(result.ID, result.Type)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"succcess": false,
			"data":     err.Error(),
		})
	}
	g.SetCookie("media", token, 3600*24, "/", "localhost", false, true)
	g.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

func (h *cmsHandler) Login() {}
