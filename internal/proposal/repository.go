package proposal

import (
	"context"
	"gh6-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Repository adalah antarmuka untuk operasi database proposal.
type Repository interface {
	Create(ctx context.Context, proposal *domain.Proposal) error
	FindByProjectID(ctx context.Context, projectID string) ([]domain.Proposal, error)
}

type repository struct {
	db *mongo.Collection
}

// NewRepository membuat repository proposal baru.
func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db.Collection("proposals"),
	}
}

func (r *repository) Create(ctx context.Context, proposal *domain.Proposal) error {
	proposal.CreatedAt = time.Now()
	proposal.UpdatedAt = time.Now()
	_, err := r.db.InsertOne(ctx, proposal)
	return err
}

func (r *repository) FindByProjectID(ctx context.Context, projectID string) ([]domain.Proposal, error) {
	var proposals []domain.Proposal
	cursor, err := r.db.Find(ctx, bson.M{"project_id": projectID, "deleted_at": nil})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &proposals); err != nil {
		return nil, err
	}

	return proposals, nil
}
