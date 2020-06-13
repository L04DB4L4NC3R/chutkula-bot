package repo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	Collection *mongo.Collection
}

func NewMongoRepo(c *mongo.Collection) Repository {
	return &mongoRepo{
		Collection: c,
	}
}

func (m *mongoRepo) Register(ctx context.Context, chatID string, feed string, updatedAt *time.Time) error {
	var result State
	err := m.Collection.FindOne(ctx, bson.M{"chat_id": chatID, "feed": feed}).Decode(&result)

	if err == nil {
		return errors.Errorf("Document Already Exists")
	}

	if err == mongo.ErrNoDocuments {
		_, err := m.Collection.InsertOne(ctx, bson.M{
			"chat_id":    chatID,
			"feed":       feed,
			"updated_at": updatedAt,
		})
		return err
	}
	return err
}

func (m *mongoRepo) UnRegister(ctx context.Context, chatID string, feed string) error {
	_, err := m.Collection.DeleteOne(ctx, bson.M{
		"chat_id": chatID,
		"feed":    feed,
	})

	if err == mongo.ErrNoDocuments {
		return errors.Errorf("Chat Not Found")
	}
	return err
}

func (m *mongoRepo) GetUpdatedAt(ctx context.Context, chatID string, feed string) (*time.Time, error) {

	var result State
	err := m.Collection.FindOne(ctx, bson.M{
		"chat_id": chatID,
		"feed":    feed,
	}).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result.UpdatedAt, nil
}

func (m *mongoRepo) UpdateTimeStamp(ctx context.Context, newtime *time.Time, chatID string, feed string) error {

	_, err := m.Collection.UpdateOne(ctx, bson.M{
		"chat_id": chatID,
		"feed":    feed,
	}, bson.M{"$set": bson.M{"updated_at": newtime}})

	return err
}

func (m *mongoRepo) GetUpdatedStates(ctx context.Context) ([]State, error) {
	cur, err := m.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var results []State
	for cur.Next(ctx) {
		var result State
		err := cur.Decode(&result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
