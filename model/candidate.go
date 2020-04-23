package model

import (
	"context"
	"time"
)

// Candidate model is used to store and exchange candidate information
// It is persisted in the DB in Candidates collection
type Candidate struct {
	ID				string		`json:"id" bson:"_id,omitempty"`
	FirstName		string		`json:"first_name" bson:"first_name"`
	LastName		string		`json:"last_name" bson:"last_name"`
	Email			string		`json:"email" validate:"required,email"`
	Department 		string 		`json:"department" validate:"required"`
	University 		string 		`json:"university" validate:"required"`
	Experience 		bool 		`json:"experience"`
	ApplicationDate time.Time	`json:"application_date" bson:"application_date"`
	Status 			string 		`json:"status"`
	MeetingCount 	int 		`json:"meeting_count" bson:"meeting_count"`
	NextMeeting 	*time.Time	`json:"next_meeting" bson:"next_meeting"`
	Assignee 		string 		`json:"assignee"`
}

type CandidateRepository interface {
	CreateCandidate(ctx context.Context, candidate Candidate) (Candidate, error)
	UpdateCandidate(ctx context.Context, id string, candidate Candidate) error
	ReadCandidate(ctx context.Context, id string) (Candidate, error)
	FindAllCandidates(ctx context.Context) ([]Candidate, error)
	FindCandidateByEmail(ctx context.Context, email string) (Candidate, error)
	FindAssigneesCandidates(ctx context.Context, id string) ([]Candidate, error)
	DeleteCandidate(ctx context.Context, id string) error
}

type CandidateService interface {
	CreateCandidate(ctx context.Context, candidate Candidate) (Candidate, error)
	UpdateCandidate(ctx context.Context, id string, candidate Candidate) error
	ReadCandidate(ctx context.Context, email string) (Candidate, error)
	FindAllCandidates(ctx context.Context) ([]Candidate, error)
	FindCandidateByEmail(ctx context.Context, id string) (Candidate, error)
	FindAssigneesCandidates(ctx context.Context, id string) ([]Candidate, error)
	DeleteCandidate(ctx context.Context, id string) error
	DenyCandidate(ctx context.Context, id string) error
	AcceptCandidate(ctx context.Context, id string) error
	ArrangeMeeting(ctx context.Context, id string, nextMeetingTime *time.Time) error
	CompleteMeeting(ctx context.Context, id string) error
}
