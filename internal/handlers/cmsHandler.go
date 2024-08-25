package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/news-system/internal/auth"
	"github.com/hritesh04/news-system/internal/core/dto"
	"github.com/hritesh04/news-system/internal/core/ports"
	"github.com/hritesh04/news-system/internal/helper"
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
		helper.ReturnFailed(g, http.StatusBadRequest, err.Error())
	}
	result, err := h.cmsService.CreateUser(user)
	if err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err.Error())
	}
	token, err := auth.GenerateToken(result.ID, result.Type)
	if err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err.Error())
	}
	g.SetCookie("media", token, 3600*24, "/", "localhost", false, true)
	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *cmsHandler) Login(g *gin.Context) {
	var credentials dto.LogInRequest
	if err := g.ShouldBindJSON(&credentials); err != nil {
		helper.ReturnFailed(g, http.StatusBadRequest, err.Error())
	}
	result, err := h.cmsService.SignInUser(credentials)
	if err != nil {
		helper.ReturnFailed(g, http.StatusUnauthorized, err)
	}
	token, err := auth.GenerateToken(result.ID, result.Type)
	if err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	g.SetCookie("media", token, 3600*24, "/", "localhost", false, true)
	helper.ReturnSuccess(g, http.StatusOK, result)
}
