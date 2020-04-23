package model

import (
	"errors"
)

var (
	ErrAssigneeDoesNotExist   = errors.New("assignee does not exist")
	ErrCandidateDoesNotExist  = errors.New("candidate does not exist")
	ErrCandidateAlreadyExists = errors.New("candidate already exist")
	ErrMeetingCountNotEnough  = errors.New("candidates cannot be accepted before the completion of 4 meetings")
	ErrArrangedMeetingDoesNotExist  = errors.New("current candidate does not have any arranged meetings")
	ErrDepartmentDoesNotExist  = errors.New("department does not exist")
)
