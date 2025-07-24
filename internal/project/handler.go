package project

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler menangani request HTTP untuk proyek.
type Handler struct {
	service Service
}

// NewHandler membuat handler proyek baru.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterRoutes mengatur rute untuk proyek.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/projects", h.CreateProject)
	router.GET("/projects/:id", h.GetProjectByID)
	router.GET("/projects", h.GetAllProjects)
}

func (h *Handler) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.service.CreateProject(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, ErrInvalidBudget) || errors.Is(err, ErrGovernmentWallet) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *Handler) GetProjectByID(c *gin.Context) {
	id := c.Param("id")

	project, err := h.service.GetProjectByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *Handler) GetAllProjects(c *gin.Context) {
	projects, err := h.service.GetAllProjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}
