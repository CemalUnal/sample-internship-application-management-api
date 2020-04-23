package repository

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mongodbCandidateRepository struct {
	collection *mongo.Collection
}

// MongoDBCandidateRepository will create an implementation of Candidate Repository with MongoDB
func MongoDBCandidateRepository(collection *mongo.Collection) model.CandidateRepository {
	return &mongodbCandidateRepository {
		collection: collection,
	}
}

func (repository *mongodbCandidateRepository) CreateCandidate(ctx context.Context, candidate model.Candidate) (model.Candidate, error) {
	_, err := repository.collection.InsertOne(ctx, candidate)
	if err != nil {
		log.Println(err)
	}

	return candidate, err
}

func (repository *mongodbCandidateRepository) UpdateCandidate(ctx context.Context, id string, candidate model.Candidate) error {
	collection := repository.collection
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", candidate},
		},
	)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (repository *mongodbCandidateRepository) ReadCandidate(ctx context.Context, id string) (model.Candidate, error) {
	var candidate model.Candidate
	err := repository.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&candidate)
	if err != nil {
		log.Println(err)
	}

	return candidate, err
}

func (repository *mongodbCandidateRepository) FindAllCandidates(ctx context.Context) ([]model.Candidate, error) {
	var candidates []model.Candidate
	cursor, err := repository.collection.Find(ctx, bson.D{})
	err = cursor.All(ctx, &candidates)
	if err != nil {
		log.Println(err)
	}

	return candidates, err
}

func (repository *mongodbCandidateRepository) FindCandidateByEmail(ctx context.Context, email string) (model.Candidate, error) {
	var candidate model.Candidate
	err := repository.collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&candidate)
	if err != nil {
		log.Println(err)
	}

	return candidate, err
}

func (repository *mongodbCandidateRepository) FindAssigneesCandidates(ctx context.Context, id string) ([]model.Candidate, error) {
	var candidates []model.Candidate
	cursor, err := repository.collection.Find(ctx, bson.D{{"assignee", id}})
	err = cursor.All(ctx, &candidates)
	if err != nil {
		log.Println(err)
	}

	return candidates, err
}

func (repository *mongodbCandidateRepository) DeleteCandidate(ctx context.Context, id string) error {
	_, err := repository.collection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		log.Println(err)
	}

	return err
}
