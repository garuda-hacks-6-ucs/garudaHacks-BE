package profile

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for profiles.
type Handler struct {
	service Service
}

// NewHandler creates a new profile handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterRoutes sets up the routes for profiles.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/profiles", h.CreateProfile)
	router.GET("/profiles/:walletAddress", h.GetProfileByWallet)
}

func (h *Handler) CreateProfile(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, ErrProfileExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrInvalidDetails) || errors.Is(err, ErrInvalidRole) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

func (h *Handler) GetProfileByWallet(c *gin.Context) {
	walletAddress := c.Param("walletAddress")

	profile, err := h.service.GetByWallet(c.Request.Context(), walletAddress)
	if err != nil {
		if errors.Is(err, ErrProfileNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
