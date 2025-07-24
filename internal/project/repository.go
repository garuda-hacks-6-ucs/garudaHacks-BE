package project

import (
	"context"
	"gh6-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Repository adalah antarmuka untuk operasi database proyek.
type Repository interface {
	Create(ctx context.Context, project *domain.Project) error
	FindByID(ctx context.Context, id string) (*domain.Project, error)
	FindAll(ctx context.Context) ([]domain.Project, error)
}

type repository struct {
	db *mongo.Collection
}

// NewRepository membuat repository proyek baru.
func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db.Collection("projects"),
	}
}

func (r *repository) Create(ctx context.Context, project *domain.Project) error {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	_, err := r.db.InsertOne(ctx, project)
	return err
}

func (r *repository) FindByID(ctx context.Context, id string) (*domain.Project, error) {
	var project domain.Project
	err := r.db.FindOne(ctx, bson.M{"_id": id, "deleted_at": nil}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func (r *repository) FindAll(ctx context.Context) ([]domain.Project, error) {
	var projects []domain.Project
	cursor, err := r.db.Find(ctx, bson.M{"deleted_at": nil})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
