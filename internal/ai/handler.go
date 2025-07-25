package ai

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler menangani request HTTP untuk fitur AI.
type Handler struct {
	service Service
}

// NewHandler membuat handler AI baru.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterRoutes mengatur rute untuk fitur AI.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/ai/summarize-pdf", h.SummarizePDF)
}

// SummarizePDF menangani upload file PDF dan mengembalikan rangkumannya.
func (h *Handler) SummarizePDF(c *gin.Context) {
	file, err := c.FormFile("proposal_pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File 'proposal_pdf' is required"})
		return
	}

	// Memeriksa tipe file
	if file.Header.Get("Content-Type") != "application/pdf" {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "File must be a PDF"})
		return
	}

	summary, err := h.service.SummarizePDF(c.Request.Context(), file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to summarize PDF", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}
