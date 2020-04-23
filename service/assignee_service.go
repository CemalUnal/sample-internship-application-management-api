package service

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type assigneeService struct {
	assigneeRepository model.AssigneeRepository
}

// AssigneeService will create an implementation of AssigneeService interface
func AssigneeService(assigneeRepository model.AssigneeRepository) model.AssigneeService {
	return &assigneeService{
		assigneeRepository: assigneeRepository,
	}
}

func (service *assigneeService) CreateAssignee(ctx context.Context, assignee model.Assignee) (model.Assignee, error) {
	assignee.ID = primitive.NewObjectID().Hex()

	return service.assigneeRepository.CreateAssignee(ctx, assignee)
}

func (service *assigneeService) FindAllAssignees(ctx context.Context) ([]model.Assignee, error) {
	return service.assigneeRepository.FindAllAssignees(ctx)
}

func (service *assigneeService) FindAllAssigneesByDepartment(ctx context.Context, department string) ([]model.Assignee, error) {
	return service.assigneeRepository.FindAllAssigneesByDepartment(ctx, department)
}

func (service *assigneeService) FindAssigneeIDByName(ctx context.Context, name string) string {
	id, _ := service.assigneeRepository.FindAssigneeIDByName(ctx, name)

	return id
}
