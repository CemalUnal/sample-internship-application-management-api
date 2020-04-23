package mocks

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/stretchr/testify/mock"
)

type AssigneeRepository struct {
	mock.Mock
}

func (a *AssigneeRepository) CreateAssignee(ctx context.Context, assignee model.Assignee) (model.Assignee, error) {
	ret := a.Called(ctx, assignee)

	var r0 model.Assignee
	if rf, ok := ret.Get(0).(func(context.Context, model.Assignee) model.Assignee); ok {
		r0 = rf(ctx, assignee)
	} else {
		r0 = ret.Get(0).(model.Assignee)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Assignee) error); ok {
		r1 = rf(ctx, assignee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (a *AssigneeRepository) ReadAssignee(ctx context.Context, id string) (model.Assignee, error) {
	ret := a.Called(ctx, id)

	var r0 model.Assignee
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Assignee); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Assignee)
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

func (a *AssigneeRepository) FindAllAssignees(ctx context.Context) ([]model.Assignee, error) {
	ret := a.Called(ctx)

	var r0 []model.Assignee
	if rf, ok := ret.Get(0).(func(context.Context) []model.Assignee); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Assignee)
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

func (a *AssigneeRepository) FindAssigneeIDByName(ctx context.Context, name string) (string, error) {
	ret := a.Called(ctx, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (a *AssigneeRepository) FindAllAssigneesByDepartment(ctx context.Context, department string) ([]model.Assignee, error) {
	ret := a.Called(ctx)

	var r0 []model.Assignee
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Assignee); ok {
		r0 = rf(ctx, department)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Assignee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, department)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (a *AssigneeRepository) FindOneAssigneeByDepartment(ctx context.Context, department string) (model.Assignee, error) {
	ret := a.Called(ctx, department)

	var r0 model.Assignee
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Assignee); ok {
		r0 = rf(ctx, department)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Assignee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, department)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
