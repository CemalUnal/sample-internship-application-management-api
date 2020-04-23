package service

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/cemalunal/sample-internship-management-api/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCandidateService_CreateCandidate(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Experience: false,
		Assignee: "123123123123",
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("FindCandidateByEmail", mock.Anything,
			mock.AnythingOfType("string")).Return(model.Candidate{}, model.ErrCandidateDoesNotExist).Once()

		mockCandidateRepository.On("CreateCandidate", mock.Anything,
			mock.AnythingOfType("model.Candidate")).Return(model.Candidate{}, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		savedCandidate, err := cService.CreateCandidate(context.TODO(), mockCandidate)

		assert.NoError(t, err)
		assert.NotNil(t, savedCandidate)
		assert.NotNil(t, savedCandidate.ID)
		mockCandidateRepository.AssertExpectations(t)
	})

	t.Run("candidate-already-exists", func(t *testing.T) {
		existingCandidate := mockCandidate
		mockCandidateRepository.On("FindCandidateByEmail", mock.Anything,
			mock.AnythingOfType("string")).Return(existingCandidate, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		_, err := cService.CreateCandidate(context.TODO(), mockCandidate)

		assert.Equal(t, err, model.ErrCandidateAlreadyExists)
		mockCandidateRepository.AssertExpectations(t)
	})
}

func TestCandidateService_UpdateCandidate(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		ID: "123asd123",
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Experience: false,
		Assignee: "123123123123",
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("UpdateCandidate", mock.Anything, mock.AnythingOfType("string"), mockCandidate).Once().Return(nil)

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		err := cService.UpdateCandidate(context.TODO(), mockCandidate.ID, mockCandidate)
		assert.NoError(t, err)
		mockCandidateRepository.AssertExpectations(t)
	})

}

func TestCandidateService_ReadCandidate(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Experience: false,
		Assignee: "123123123123",
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(mockCandidate, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)

		foundCandidate, err := cService.ReadCandidate(context.TODO(), mockCandidate.ID)

		assert.Equal(t, mockCandidate, foundCandidate)
		assert.NoError(t, err)
		mockCandidateRepository.AssertExpectations(t)
	})
}

func TestCandidateService_FindAllCandidates(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidateArray := []model.Candidate {
		{
			FirstName:  "testFN",
			LastName:   "testLN",
			Email:      "t@t.com",
			Department: "CS",
			University: "HU",
			Experience: false,
			Assignee:   "123123123123",
		},
		{
			FirstName:  "testFN2",
			LastName:   "testLN2",
			Email:      "t2@t2.com",
			Department: "CS",
			University: "HU",
			Experience: false,
			Assignee:   "123123123123",
		},
		{
			FirstName:  "testFN3",
			LastName:   "testLN3",
			Email:      "t3@t3.com",
			Department: "CS",
			University: "HU",
			Experience: false,
			Assignee:   "123123123123",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("FindAllCandidates", mock.Anything).Return(mockCandidateArray, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		candidateArray, err := cService.FindAllCandidates(context.TODO())

		assert.NoError(t, err)
		assert.NotNil(t, candidateArray)
		assert.Equal(t, mockCandidateArray, candidateArray)
		mockCandidateRepository.AssertExpectations(t)
	})
}

func TestCandidateService_FindCandidateByEmail(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Experience: false,
		Assignee: "123123123123",
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("FindCandidateByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockCandidate, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		foundCandidate, err := cService.FindCandidateByEmail(context.TODO(), mockCandidate.Email)

		assert.Equal(t, mockCandidate, foundCandidate)
		assert.NoError(t, err)
		mockCandidateRepository.AssertExpectations(t)
	})
}

func TestCandidateService_FindAssigneesCandidates(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockAssignee := model.Assignee{
		ID: "asd123dsa",
		Name: "A1",
		Department: "CS",
	}
	mockCandidateArray := []model.Candidate {
		{
			FirstName:  "testFN",
			LastName:   "testLN",
			Email:      "t@t.com",
			Department: "CS",
			University: "HU",
			Experience: false,
			Assignee:   "123123123123",
		},
		{
			FirstName:  "testFN2",
			LastName:   "testLN2",
			Email:      "t2@t2.com",
			Department: "CS",
			University: "HU",
			Experience: false,
			Assignee:   "123123123123",
		},
		{
			FirstName:  "testFN3",
			LastName:   "testLN3",
			Email:      "t3@t3.com",
			Department: "CS",
			University: "HU",
			Experience: false,
			Assignee:   "123123123123",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("FindAssigneesCandidates", mock.Anything,
			mock.AnythingOfType("string")).Return(mockCandidateArray, nil).Once()

		mockAssigneeRepository.On("ReadAssignee", mock.Anything,
			mock.AnythingOfType("string")).Return(mockAssignee, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		candidateArray, err := cService.FindAssigneesCandidates(context.TODO(), mockAssignee.ID)

		assert.NoError(t, err)
		assert.NotNil(t, candidateArray)
		assert.Equal(t, mockCandidateArray, candidateArray)
		mockCandidateRepository.AssertExpectations(t)
		mockAssigneeRepository.AssertExpectations(t)
	})

	t.Run("assignee-does-not-exist", func(t *testing.T) {
		mockAssigneeRepository.On("ReadAssignee", mock.Anything,
			mock.AnythingOfType("string")).Return(model.Assignee{}, model.ErrAssigneeDoesNotExist).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		_, err := cService.FindAssigneesCandidates(context.TODO(), mockAssignee.ID)

		assert.Equal(t, err, model.ErrAssigneeDoesNotExist)
		mockCandidateRepository.AssertExpectations(t)
		mockAssigneeRepository.AssertExpectations(t)
	})
}

func TestCandidateService_DeleteCandidate(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Experience: false,
		Assignee: "123123123123",
	}

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(mockCandidate, nil).Once()
		mockCandidateRepository.On("DeleteCandidate", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)

		err := cService.DeleteCandidate(context.TODO(), mockCandidate.ID)

		assert.NoError(t, err)
		mockCandidateRepository.AssertExpectations(t)
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(model.Candidate{}, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)

		err := cService.DeleteCandidate(context.TODO(), mockCandidate.ID)

		assert.Equal(t, err, model.ErrCandidateDoesNotExist)
		mockCandidateRepository.AssertExpectations(t)
	})
}

func TestCandidateService_DenyCandidate(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		ID: "123asd123",
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Status: model.InProgress,
		Experience: false,
		Assignee: "123123123123",
	}
	mockDeniedCandidate := mockCandidate
	mockDeniedCandidate.Status = model.Denied

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(mockCandidate, nil).Once()
		mockCandidateRepository.On("UpdateCandidate", mock.Anything, mockCandidate.ID, mockDeniedCandidate).Once().Return(nil)

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		err := cService.DenyCandidate(context.TODO(), mockCandidate.ID)

		assert.NoError(t, err)
		mockCandidateRepository.AssertExpectations(t)
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(model.Candidate{}, nil).Once()
		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)

		err := cService.DenyCandidate(context.TODO(), mockCandidate.ID)

		assert.Equal(t, err, model.ErrCandidateDoesNotExist)
		mockCandidateRepository.AssertExpectations(t)
	})
}

func TestCandidateService_AcceptCandidate(t *testing.T) {
	mockCandidateRepository := new(mocks.CandidateRepository)
	mockAssigneeRepository := new(mocks.AssigneeRepository)
	mockCandidate := model.Candidate{
		ID: "123asd123",
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: "CS",
		University: "HU",
		Status: model.InProgress,
		Experience: false,
		MeetingCount: 4,
		Assignee: "123123123123",
	}
	mockAcceptedCandidate := mockCandidate
	mockLowMeetingCountCandidate := mockCandidate
	mockAcceptedCandidate.Status = model.Accepted
	mockLowMeetingCountCandidate.MeetingCount = 2

	t.Run("success", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(mockCandidate, nil).Once()
		mockCandidateRepository.On("UpdateCandidate", mock.Anything, mockCandidate.ID, mockAcceptedCandidate).Once().Return(nil)

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		err := cService.AcceptCandidate(context.TODO(), mockCandidate.ID)

		assert.NoError(t, err)
		mockCandidateRepository.AssertExpectations(t)
	})

	t.Run("meeting-count-is-below-four", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(mockLowMeetingCountCandidate, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		err := cService.AcceptCandidate(context.TODO(), mockLowMeetingCountCandidate.ID)

		assert.Equal(t, err, model.ErrMeetingCountNotEnough)
		mockCandidateRepository.AssertExpectations(t)
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		mockCandidateRepository.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(model.Candidate{}, nil).Once()

		cService := CandidateService(mockCandidateRepository, mockAssigneeRepository)
		err := cService.AcceptCandidate(context.TODO(), mockCandidate.ID)

		assert.Equal(t, err, model.ErrCandidateDoesNotExist)
		mockCandidateRepository.AssertExpectations(t)
	})
}
