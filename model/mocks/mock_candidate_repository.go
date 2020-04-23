package mocks

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/stretchr/testify/mock"
)

type CandidateRepository struct {
	mock.Mock
}

func (c *CandidateRepository) CreateCandidate(ctx context.Context, candidate model.Candidate) (model.Candidate, error) {
	ret := c.Called(ctx, candidate)

	var r0 model.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, model.Candidate) model.Candidate); ok {
		r0 = rf(ctx, candidate)
	} else {
		r0 = ret.Get(0).(model.Candidate)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Candidate) error); ok {
		r1 = rf(ctx, candidate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *CandidateRepository) UpdateCandidate(ctx context.Context, id string, candidate model.Candidate) error {
	ret := c.Called(ctx, id, candidate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Candidate) error); ok {
		r0 = rf(ctx, id, candidate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (c *CandidateRepository) ReadCandidate(ctx context.Context, id string) (model.Candidate, error) {
	ret := c.Called(ctx, id)

	var r0 model.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Candidate); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *CandidateRepository) FindAllCandidates(ctx context.Context) ([]model.Candidate, error) {
	ret := c.Called(ctx)

	var r0 []model.Candidate
	if rf, ok := ret.Get(0).(func(context.Context) []model.Candidate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *CandidateRepository) FindCandidateByEmail(ctx context.Context, email string) (model.Candidate, error) {
	ret := c.Called(ctx, email)

	var r0 model.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Candidate); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *CandidateRepository) FindAssigneesCandidates (ctx context.Context, id string) ([]model.Candidate, error) {
	ret := c.Called(ctx, id)

	var r0 []model.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Candidate); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *CandidateRepository) DeleteCandidate(ctx context.Context, id string) error {
	ret := c.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
