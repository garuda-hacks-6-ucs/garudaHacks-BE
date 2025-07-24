package profile

import (
	"context"
	"encoding/json"
	"errors"
	"gh6-2/internal/domain"
	"time"
)

var (
	ErrProfileExists   = errors.New("profile with this wallet address already exists")
	ErrProfileNotFound = errors.New("profile not found")
	ErrInvalidRole     = errors.New("invalid role provided")
	ErrInvalidDetails  = errors.New("details object does not match the role")
)

// Service is the interface for profile business logic.
type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*domain.Profile, error)
	GetByWallet(ctx context.Context, walletAddress string) (*domain.Profile, error)
}

type service struct {
	repo Repository
}

// NewService creates a new profile service.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// RegisterRequest defines the payload for creating a new profile.
type RegisterRequest struct {
	WalletAddress string          `json:"wallet_address" binding:"required"`
	Role          string          `json:"role" binding:"required,oneof=GOVERNMENT VENDOR CITIZEN"`
	Details       json.RawMessage `json:"details" binding:"required"`
}

func (s *service) Register(ctx context.Context, req RegisterRequest) (*domain.Profile, error) {
	existing, err := s.repo.FindByWalletAddress(ctx, req.WalletAddress)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrProfileExists
	}

	profile := &domain.Profile{
		WalletAddress: req.WalletAddress,
		Role:          req.Role,
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Validate and unmarshal details based on role
	switch req.Role {
	case "GOVERNMENT":
		var details domain.GovernmentDetails
		if err := json.Unmarshal(req.Details, &details); err != nil {
			return nil, ErrInvalidDetails
		}
		profile.Details = details
	case "VENDOR":
		var details domain.VendorDetails
		if err := json.Unmarshal(req.Details, &details); err != nil {
			return nil, ErrInvalidDetails
		}
		profile.Details = details
	case "CITIZEN":
		var details domain.CitizenDetails
		if err := json.Unmarshal(req.Details, &details); err != nil {
			return nil, ErrInvalidDetails
		}
		profile.Details = details
	default:
		return nil, ErrInvalidRole
	}

	if err := s.repo.Create(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *service) GetByWallet(ctx context.Context, walletAddress string) (*domain.Profile, error) {
	profile, err := s.repo.FindByWalletAddress(ctx, walletAddress)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, ErrProfileNotFound
	}
	return profile, nil
}
