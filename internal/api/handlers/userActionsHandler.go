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

type userActionHandler struct {
	logger            *slog.Logger
	userActionService ports.UserActionService
}

func SetupUserActionRoutes(logger *slog.Logger, rh rest.RestHandler) {
	userActionRepo := repositories.NewUserActionRepository(rh.DB)
	svc := services.NewUserActionService(userActionRepo)

	handler := userActionHandler{
		logger:            logger,
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
		h.logger.Error("request body parsing", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	user := g.GetHeader("userID")
	userID, err := strconv.ParseUint(user, 10, 32)
	if err != nil {
		h.logger.Error("parsing userID from headers", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	comment.UserID = uint(userID)
	result, err := h.userActionService.AddComment(comment)
	if err != nil {
		h.logger.Error("creating comment", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *userActionHandler) RemoveComment(g *gin.Context) {
	commentID := g.Param("commentID")
	if err := h.userActionService.RemoveComment(commentID); err != nil {
		h.logger.Error("removing comment", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, "comment deleted sucessfully")
}

func (h *userActionHandler) Subscribe(g *gin.Context) {
	var subscription dto.SubscriptionRequest
	if err := g.ShouldBindJSON(&subscription); err != nil {
		h.logger.Error("request body parsing", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	user := g.GetHeader("userID")
	userId, err := strconv.ParseUint(user, 10, 32)
	if err != nil {
		h.logger.Error("parsing userID from headers", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	subscription.UserID = uint(userId)
	result, err := h.userActionService.Subscribe(subscription)
	if err != nil {
		h.logger.Error("subscribing", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}

	helper.ReturnSuccess(g, http.StatusOK, result)
}

func (h *userActionHandler) UnSubscribe(g *gin.Context) {
	if err := h.userActionService.UnSubscribe(g.Param("subscriptionId")); err != nil {
		h.logger.Error("unsubscribing", "error", err)
		helper.ReturnFailed(g, http.StatusInternalServerError, err)
	}
	helper.ReturnSuccess(g, http.StatusOK, "unsubscribed sucessfully")
}
