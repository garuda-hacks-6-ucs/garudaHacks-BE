package server

import (
	"fmt"
	"gh6-2/internal/profile"
	"gh6-2/internal/project"
	"gh6-2/internal/proposal"
	"github.com/gin-gonic/gin"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	router *gin.Engine
	port   string
}

// New creates a new HTTP server.
func New(port string) *Server {
	router := gin.Default()
	return &Server{
		router: router,
		port:   port,
	}
}

// RegisterHandlers sets up all the API routes.
func (s *Server) RegisterHandlers(
	profileHandler *profile.Handler,
	projectHandler *project.Handler,
	proposalHandler *proposal.Handler,
) {
	api := s.router.Group("/api/v1")

	// Health check
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// Register handlers
	profileHandler.RegisterRoutes(api)
	projectHandler.RegisterRoutes(api)
	proposalHandler.RegisterRoutes(api)
}

// Run starts the HTTP server.
func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.port)
	fmt.Printf("Server running on port %s\n", s.port)
	return s.router.Run(addr)
}
