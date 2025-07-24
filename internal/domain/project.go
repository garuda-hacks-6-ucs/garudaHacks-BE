package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project represents a government tender.
type Project struct {
	ID                   string               `bson:"_id"`
	GovernmentWallet     string               `bson:"government_wallet"`
	ProjectName          string               `bson:"project_name"`
	Description          string               `bson:"description"`
	Images               []string             `bson:"images"`
	BudgetWei            primitive.Decimal128 `bson:"budget_wei"`
	Status               string               `bson:"status"`
	SmartContractAddress string               `bson:"smart_contract_address,omitempty"`
	WinningProposalID    *string              `bson:"winning_proposal_id,omitempty"`
	ProposalDeadline     time.Time            `bson:"proposal_deadline"`
	VotingDeadline       time.Time            `bson:"voting_deadline"`
	CreatedAt            time.Time            `bson:"created_at"`
	UpdatedAt            time.Time            `bson:"updated_at"`
	DeletedAt            *time.Time           `bson:"deleted_at,omitempty"`
}
