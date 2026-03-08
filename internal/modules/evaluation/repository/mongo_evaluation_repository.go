package repository

import (
	"context"
	"time"

	"github.com/rdl/core/internal/modules/evaluation/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoEvaluationRepository struct {
	collection *mongo.Collection
}

func NewMongoEvaluationRepository(db *mongo.Database) EvaluationRepository {
	return &mongoEvaluationRepository{
		collection: db.Collection("evaluations"),
	}
}

func (r *mongoEvaluationRepository) Create(ctx context.Context, e *domain.Evaluation) error {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, e)
	return err
}

func (r *mongoEvaluationRepository) FindByID(ctx context.Context, id string) (*domain.Evaluation, error) {
	var e domain.Evaluation
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *mongoEvaluationRepository) List(ctx context.Context, filter domain.ListEvaluationsFilter) ([]*domain.Evaluation, error) {
	query := bson.M{}
	if filter.ProjectID != "" {
		query["project_id"] = filter.ProjectID
	}
	if filter.EvaluatorID != "" {
		query["evaluator_id"] = filter.EvaluatorID
	}

	skip := int64((filter.Page - 1) * filter.Limit)
	limit := int64(filter.Limit)
	opts := options.Find().SetSkip(skip).SetLimit(limit)

	cur, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var items []*domain.Evaluation
	if err := cur.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *mongoEvaluationRepository) Update(ctx context.Context, e *domain.Evaluation) error {
	e.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": e.ID}, bson.M{"$set": e})
	return err
}

func (r *mongoEvaluationRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
