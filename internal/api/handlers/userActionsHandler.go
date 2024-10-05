package handlers

import (
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

type userActionHandler struct {
	userActionService ports.UserActionService
}

func SetupUserActionRoutes(rh rest.RestHandler) {
	userActionRepo := repositories.NewUserActionRepository(rh.DB)
	svc := services.NewUserActionService(userActionRepo)

	handler := userActionHandler{
		userActionService: svc,
	}
	commentGroup := rh.Router.Group("/comment")
	subscriptionGroup := rh.Router.Group("/subscription")
	commentGroup.POST("/", handler.AddComment)
	commentGroup.DELETE("/:commentId", handler.RemoveComment)
	subscriptionGroup.POST("/", handler.Subscribe)
	subscriptionGroup.DELETE("/:subscriptionId", handler.UnSubscribe)
}

func (h *userActionHandler) AddComment(g *gin.Context) {
	var comment dto.Comment
	if err := g.ShouldBindJSON(&comment); err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	user := g.GetHeader("userID")
	userID, err := strconv.ParseUint(user, 10, 32)
	if err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	comment.UserID = uint(userID)
	result, err := h.userActionService.AddComment(comment)
	if err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *userActionHandler) RemoveComment(g *gin.Context) {
	commentID := g.Param("commentID")
	if err := h.userActionService.RemoveComment(commentID); err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, "comment deleted sucessfully")
}

func (h *userActionHandler) Subscribe(g *gin.Context) {
	var subscription dto.SubscriptionRequest
	if err := g.ShouldBindJSON(&subscription); err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	user := g.GetHeader("userID")
	userId, err := strconv.ParseUint(user, 10, 32)
	if err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	subscription.UserID = uint(userId)
	result, err := h.userActionService.Subscribe(subscription)
	if err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}

	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *userActionHandler) UnSubscribe(g *gin.Context) {
	if err := h.userActionService.UnSubscribe(g.Param("subscriptionId")); err != nil {
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, "unsubscripted sucessfully")
}
