package profile

import (
	"context"
	"gh6-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Repository is the interface for database operations.
type Repository interface {
	Create(ctx context.Context, profile *domain.Profile) error
	FindByWalletAddress(ctx context.Context, walletAddress string) (*domain.Profile, error)
}

type repository struct {
	db *mongo.Collection
}

// NewRepository creates a new profile repository.
func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db.Collection("profiles"),
	}
}

func (r *repository) Create(ctx context.Context, profile *domain.Profile) error {
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()
	_, err := r.db.InsertOne(ctx, profile)
	return err
}

func (r *repository) FindByWalletAddress(ctx context.Context, walletAddress string) (*domain.Profile, error) {
	var profile domain.Profile
	filter := bson.M{"wallet_address": walletAddress, "deleted_at": nil}
	err := r.db.FindOne(ctx, filter).Decode(&profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil, nil if not found
		}
		return nil, err
	}
	return &profile, nil
}
