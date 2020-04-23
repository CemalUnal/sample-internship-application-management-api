package service

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/cemalunal/sample-internship-management-api/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAssigneeService_CreateAssignee(t *testing.T) {
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockAssignee := model.Assignee{
		Name: "FN",
		Department: "CS",
	}

	t.Run("success", func(t *testing.T) {
		mockAssigneeRepository.On("CreateAssignee", mock.Anything,
			mock.AnythingOfType("model.Assignee")).Return(model.Assignee{}, nil).Once()

		aService := AssigneeService(mockAssigneeRepository)
		savedAssignee, err := aService.CreateAssignee(context.TODO(), mockAssignee)

		assert.NoError(t, err)
		assert.NotNil(t, savedAssignee)
		assert.NotNil(t, savedAssignee.ID)
		mockAssigneeRepository.AssertExpectations(t)
	})
}

func TestAssigneeService_FindAllAssignees(t *testing.T) {
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockAssigneeArray := []model.Assignee {
		{
			Name: "test1",
			Department: model.Design,
		},
		{
			Name: "test2",
			Department: model.Development,
		},
		{
			Name: "test3",
			Department: model.Development,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockAssigneeRepository.On("FindAllAssignees", mock.Anything).Return(mockAssigneeArray, nil).Once()

		aService := AssigneeService(mockAssigneeRepository)
		assigneeArray, err := aService.FindAllAssignees(context.TODO())

		assert.NoError(t, err)
		assert.NotNil(t, assigneeArray)
		assert.Equal(t, mockAssigneeArray, assigneeArray)
		mockAssigneeRepository.AssertExpectations(t)
	})
}

func TestAssigneeService_FindAllAssigneesByDepartment(t *testing.T) {
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockAssigneeArray := []model.Assignee {
		{
			Name: "test2",
			Department: model.Development,
		},
		{
			Name: "test3",
			Department: model.Development,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockAssigneeRepository.On("FindAllAssigneesByDepartment", mock.Anything).Return(mockAssigneeArray, nil).Once()

		aService := AssigneeService(mockAssigneeRepository)
		assigneeArray, err := aService.FindAllAssigneesByDepartment(context.TODO(), model.Development)

		assert.NoError(t, err)
		assert.NotNil(t, assigneeArray)
		assert.Equal(t, mockAssigneeArray, assigneeArray)
		mockAssigneeRepository.AssertExpectations(t)
	})
}

func TestAssigneeService_FindAssigneeIDByName(t *testing.T) {
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockAssignee := model.Assignee {
		ID: "asd123qwe",
		Name: "test2",
		Department: model.Development,
	}

	t.Run("success", func(t *testing.T) {
		mockAssigneeRepository.On("FindAssigneeIDByName", mock.Anything, mock.AnythingOfType("string")).Return(mockAssignee.ID, nil).Once()

		aService := AssigneeService(mockAssigneeRepository)
		foundId := aService.FindAssigneeIDByName(context.TODO(), mockAssignee.Name)

		assert.NotNil(t, foundId)
		assert.Equal(t, mockAssignee.ID, foundId)
		mockAssigneeRepository.AssertExpectations(t)
	})
}
