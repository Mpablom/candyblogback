package work

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func newRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("works"),
	}
}

func (r *Repository) GetAllWorks() ([]Work, error) {
	var works []Work
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &works); err != nil {
		return nil, err
	}
	return works, nil
}

func (r *Repository) GetWork(id primitive.ObjectID) (*Work, error) {
	var work Work
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&work)
	if err != nil {
		return nil, err
	}
	return &work, nil
}

func (r *Repository) CreateWork(work *Work) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, work)
	return err
}

func (r *Repository) UpdateWork(work *Work) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": work.ID}
	update := bson.M{"$set": work}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *Repository) DeleteWork(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
