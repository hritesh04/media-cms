package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *cmsHandler) Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "msg")
}
