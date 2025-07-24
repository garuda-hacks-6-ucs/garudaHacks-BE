package main

import (
	"context"
	"gh6-2/internal/config"
	"gh6-2/internal/platform/database"
	"gh6-2/internal/platform/server"
	"gh6-2/internal/profile"
	"gh6-2/internal/project"
	"gh6-2/internal/proposal"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a new context
	ctx := context.Background()

	// Connect to MongoDB
	db, err := database.NewMongoConnection(ctx, cfg.MongoURI, cfg.MongoDbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize Repositories
	profileRepo := profile.NewRepository(db)
	projectRepo := project.NewRepository(db)
	proposalRepo := proposal.NewRepository(db)

	// Initialize Services
	profileSvc := profile.NewService(profileRepo)
	projectSvc := project.NewService(projectRepo)
	proposalSvc := proposal.NewService(proposalRepo)

	// Initialize Handlers
	profileHandler := profile.NewHandler(profileSvc)
	projectHandler := project.NewHandler(projectSvc)
	proposalHandler := proposal.NewHandler(proposalSvc)

	// Initialize Server
	srv := server.New(cfg.ServerPort)
	srv.RegisterHandlers(profileHandler, projectHandler, proposalHandler)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
