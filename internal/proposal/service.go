package proposal

import (
	"context"
	"errors"
	"gh6-2/internal/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidRequest = errors.New("invalid request body")
)

// Service adalah antarmuka untuk logika bisnis proposal.
type Service interface {
	CreateProposal(ctx context.Context, req CreateProposalRequest) (*domain.Proposal, error)
	GetProposalsByProjectID(ctx context.Context, projectID string) ([]domain.Proposal, error)
}

type service struct {
	repo Repository
}

// NewService membuat service proposal baru.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// CreateProposalRequest mendefinisikan payload untuk membuat proposal baru.
type CreateProposalRequest struct {
	ProjectID          string                `json:"project_id" binding:"required"`
	VendorWallet       string                `json:"vendor_wallet" binding:"required"`
	ProposalName       string                `json:"proposal_name" binding:"required"`
	Description        string                `json:"description" binding:"required"`
	Images             []string              `json:"images"`
	RequestedBudgetWei string                `json:"requested_budget_wei" binding:"required"`
	AiSummary          string                `json:"ai_summary"`
	OnchainPayload     domain.OnchainPayload `json:"onchain_payload"`
}

func (s *service) CreateProposal(ctx context.Context, req CreateProposalRequest) (*domain.Proposal, error) {
	budget, err := primitive.ParseDecimal128(req.RequestedBudgetWei)
	if err != nil {
		return nil, errors.New("invalid requested budget value")
	}

	proposal := &domain.Proposal{
		ID:                 uuid.New().String(),
		ProjectID:          req.ProjectID,
		VendorWallet:       req.VendorWallet,
		ProposalName:       req.ProposalName,
		Description:        req.Description,
		Images:             req.Images,
		RequestedBudgetWei: budget,
		AiSummary:          req.AiSummary,
		OnchainPayload:     req.OnchainPayload,
		Status:             "SUBMITTED",
	}

	if err := s.repo.Create(ctx, proposal); err != nil {
		return nil, err
	}

	return proposal, nil
}

func (s *service) GetProposalsByProjectID(ctx context.Context, projectID string) ([]domain.Proposal, error) {
	return s.repo.FindByProjectID(ctx, projectID)
}
