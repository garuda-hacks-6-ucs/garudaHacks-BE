package server

import (
	"fmt"
	"gh6-2/internal/profile"
	"gh6-2/internal/project"
	"gh6-2/internal/proposal"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	router *gin.Engine
	port   string
}

// New creates a new HTTP server.
func New(port string) *Server {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Ini akan mengizinkan semua IP/domain
	// Anda juga bisa spesifik: config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	router.Use(cors.New(config))

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
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
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
