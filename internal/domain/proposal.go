package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OnchainPayload holds data for smart contract execution.
type OnchainPayload struct {
	Targets   []string `bson:"targets" json:"targets"`
	Values    []string `bson:"values" json:"values"`
	Calldatas []string `bson:"calldatas" json:"calldatas"`
}

// Proposal represents a vendor's submission for a project.
type Proposal struct {
	ID                 string               `bson:"_id"`
	ProjectID          string               `bson:"project_id"`
	VendorWallet       string               `bson:"vendor_wallet"`
	ProposalName       string               `bson:"proposal_name"`
	Description        string               `bson:"description"`
	Images             []string             `bson:"images"`
	RequestedBudgetWei primitive.Decimal128 `bson:"requested_budget_wei"`
	AiSummary          string               `bson:"ai_summary,omitempty"`
	OnchainPayload     OnchainPayload       `bson:"onchain_payload"`
	Status             string               `bson:"status"`
	CreatedAt          time.Time            `bson:"created_at"`
	UpdatedAt          time.Time            `bson:"updated_at"`
	DeletedAt          *time.Time           `bson:"deleted_at,omitempty"`
}
