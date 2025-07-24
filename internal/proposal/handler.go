package proposal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler menangani request HTTP untuk proposal.
type Handler struct {
	service Service
}

// NewHandler membuat handler proposal baru.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterRoutes mengatur rute untuk proposal.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/proposals", h.CreateProposal)
	router.GET("/projects/:id/proposals", h.GetProposalsByProjectID)
}

func (h *Handler) CreateProposal(c *gin.Context) {
	var req CreateProposalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	proposal, err := h.service.CreateProposal(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proposal"})
		return
	}

	c.JSON(http.StatusCreated, proposal)
}

func (h *Handler) GetProposalsByProjectID(c *gin.Context) {
	projectID := c.Param("id")

	proposals, err := h.service.GetProposalsByProjectID(c.Request.Context(), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get proposals for the project"})
		return
	}

	c.JSON(http.StatusOK, proposals)
}
