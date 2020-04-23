package model

import (
	"context"
)

// Assignee model is used to store and exchange assignee information
// It is persisted in the DB in Assignees collection
type Assignee struct {
	ID    		string 	`json:"id" bson:"_id,omitempty"`
	Name     	string 	`json:"name" validate:"required"`
	Department 	string	`json:"department" validate:"required"`
}

type AssigneeRepository interface {
	CreateAssignee(ctx context.Context, assignee Assignee) (Assignee, error)
	ReadAssignee(ctx context.Context, id string) (Assignee, error)
	FindAllAssignees(ctx context.Context) ([]Assignee, error)
	FindAssigneeIDByName(ctx context.Context, name string) (string, error)
	FindAllAssigneesByDepartment(ctx context.Context, department string) ([]Assignee, error)
	FindOneAssigneeByDepartment(ctx context.Context, department string) (Assignee, error)
}

type AssigneeService interface {
	CreateAssignee(ctx context.Context, assignee Assignee) (Assignee, error)
	FindAllAssignees(ctx context.Context) ([]Assignee, error)
	FindAllAssigneesByDepartment(ctx context.Context, department string) ([]Assignee, error)
	FindAssigneeIDByName(ctx context.Context, name string) string
}
