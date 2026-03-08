package repository

import (
	"context"
	"time"

	"github.com/rdl/core/internal/modules/project/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoProjectRepository struct {
	collection *mongo.Collection
}

func NewMongoProjectRepository(db *mongo.Database) ProjectRepository {
	return &mongoProjectRepository{
		collection: db.Collection("projects"),
	}
}

func (r *mongoProjectRepository) Create(ctx context.Context, p *domain.Project) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, p)
	return err
}

func (r *mongoProjectRepository) FindByID(ctx context.Context, id string) (*domain.Project, error) {
	var p domain.Project
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *mongoProjectRepository) List(ctx context.Context, filter domain.ListProjectsFilter) ([]*domain.Project, error) {
	query := bson.M{}
	if filter.OwnerID != "" {
		query["owner_id"] = filter.OwnerID
	}
	if filter.Status != "" {
		query["status"] = filter.Status
	}

	skip := int64((filter.Page - 1) * filter.Limit)
	limit := int64(filter.Limit)
	opts := options.Find().SetSkip(skip).SetLimit(limit)

	cur, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var projects []*domain.Project
	if err := cur.All(ctx, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *mongoProjectRepository) Update(ctx context.Context, p *domain.Project) error {
	p.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": p.ID}, bson.M{"$set": p})
	return err
}

func (r *mongoProjectRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
