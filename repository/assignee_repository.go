package repository

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongodbAssigneeRepository struct {
	collection *mongo.Collection
}

// MongoDBAssigneeRepository will create an implementation of Assignee Repository with MongoDB
func MongoDBAssigneeRepository(collection *mongo.Collection) model.AssigneeRepository {
	return &mongodbAssigneeRepository {
		collection: collection,
	}
}

func (repository *mongodbAssigneeRepository) CreateAssignee(ctx context.Context, assignee model.Assignee) (model.Assignee, error) {
	_, err := repository.collection.InsertOne(ctx, assignee)
	if err != nil {
		log.Println(err)
	}

	return assignee, err
}

func (repository *mongodbAssigneeRepository) FindAllAssignees(ctx context.Context) ([]model.Assignee, error) {
	var assignees []model.Assignee
	cursor, err := repository.collection.Find(ctx, bson.D{})
	err = cursor.All(ctx, &assignees)
	if err != nil {
		log.Println(err)
	}

	return assignees, err
}

func (repository *mongodbAssigneeRepository) ReadAssignee(ctx context.Context, id string) (model.Assignee, error) {
	var assignee model.Assignee
	err := repository.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&assignee)
	if err != nil {
		log.Println(err)
	}

	return assignee, err
}

func (repository *mongodbAssigneeRepository) FindAssigneeIDByName(ctx context.Context, name string) (string, error) {
	var assignee model.Assignee
	projection := bson.D{{"_id", 1}}

	err := repository.collection.FindOne(ctx,
		bson.D{{"name", name}},
		options.FindOne().SetProjection(projection),
	).Decode(&assignee)

	if err != nil {
		log.Println(err)
	}

	return assignee.ID, err
}

func (repository *mongodbAssigneeRepository) FindAllAssigneesByDepartment(ctx context.Context, department string) ([]model.Assignee, error) {
	var assignees []model.Assignee
	cursor, err := repository.collection.Find(ctx, bson.D{{"department", department}})
	err = cursor.All(ctx, &assignees)
	if err != nil {
		log.Println(err)
	}

	return assignees, err
}

func (repository *mongodbAssigneeRepository) FindOneAssigneeByDepartment(ctx context.Context, department string) (model.Assignee, error) {
	var assignees []model.Assignee
	matchStage := bson.D{{"$match", bson.D{{"department", department}}}}
	asd := bson.D{{"$sample", bson.D{{"size", 1}}}}

	cursor, err := repository.collection.Aggregate(ctx, mongo.Pipeline{matchStage, asd})
	err = cursor.All(ctx, &assignees)
	if err != nil {
		log.Println(err)
	}

	var assignee model.Assignee
	if len (assignees) > 0 {
		assignee = assignees[0]
	}

	return assignee, err
}
