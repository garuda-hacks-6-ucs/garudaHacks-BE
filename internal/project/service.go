package project

import (
	"context"
	"errors"
	"gh6-2/internal/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrInvalidBudget    = errors.New("invalid budget value")
	ErrGovernmentWallet = errors.New("government wallet address is required")
)

// Service adalah antarmuka untuk logika bisnis proyek.
type Service interface {
	CreateProject(ctx context.Context, req CreateProjectRequest) (*domain.Project, error)
	GetProjectByID(ctx context.Context, id string) (*domain.Project, error)
	GetAllProjects(ctx context.Context) ([]domain.Project, error)
}

type service struct {
	repo Repository
}

// NewService membuat service proyek baru.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// CreateProjectRequest mendefinisikan payload untuk membuat proyek baru.
type CreateProjectRequest struct {
	GovernmentWallet     string    `json:"government_wallet" binding:"required"`
	ProjectName          string    `json:"project_name" binding:"required"`
	Description          string    `json:"description" binding:"required"`
	Images               []string  `json:"images"`
	BudgetWei            string    `json:"budget_wei" binding:"required"`
	SmartContractAddress string    `json:"smart_contract_address"`
	ProposalDeadline     time.Time `json:"proposal_deadline" binding:"required"`
	VotingDeadline       time.Time `json:"voting_deadline" binding:"required"`
}

func (s *service) CreateProject(ctx context.Context, req CreateProjectRequest) (*domain.Project, error) {
	if req.GovernmentWallet == "" {
		return nil, ErrGovernmentWallet
	}

	budget, err := primitive.ParseDecimal128(req.BudgetWei)
	if err != nil {
		return nil, ErrInvalidBudget
	}

	project := &domain.Project{
		ID:                   uuid.New().String(),
		GovernmentWallet:     req.GovernmentWallet,
		ProjectName:          req.ProjectName,
		Description:          req.Description,
		Images:               req.Images,
		BudgetWei:            budget,
		Status:               "OPEN_FOR_PROPOSAL", // Status awal
		SmartContractAddress: req.SmartContractAddress,
		ProposalDeadline:     req.ProposalDeadline,
		VotingDeadline:       req.VotingDeadline,
	}

	if err := s.repo.Create(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *service) GetProjectByID(ctx context.Context, id string) (*domain.Project, error) {
	project, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, ErrProjectNotFound
	}
	return project, nil
}

func (s *service) GetAllProjects(ctx context.Context) ([]domain.Project, error) {
	return s.repo.FindAll(ctx)
}
