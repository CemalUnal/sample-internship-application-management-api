package service

import (
	"context"
	"github.com/cemalunal/sample-internship-management-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type candidateService struct {
	candidateRepository model.CandidateRepository
	assigneeRepository model.AssigneeRepository
}

// CandidateService will create an implementation of CandidateService interface
func CandidateService(candidateRepository model.CandidateRepository, assigneeRepository model.AssigneeRepository) model.CandidateService {
	return &candidateService{
		candidateRepository: candidateRepository,
		assigneeRepository: assigneeRepository,
	}
}

func (service *candidateService) CreateCandidate(ctx context.Context, candidate model.Candidate) (model.Candidate, error) {
	// Check candidate exists with given email, return error if exists.
	c, _ := service.FindCandidateByEmail(ctx, candidate.Email)
	if c != (model.Candidate{}) {
		log.Println(model.ErrCandidateAlreadyExists)
		return model.Candidate{}, model.ErrCandidateAlreadyExists
	}

	candidate.ID = primitive.NewObjectID().Hex()
	candidate.Status = model.Pending
	candidate.MeetingCount = 0
	candidate.NextMeeting = nil
	candidate.ApplicationDate = time.Now()

	return service.candidateRepository.CreateCandidate(ctx, candidate)
}

func (service *candidateService) UpdateCandidate(ctx context.Context, id string, candidate model.Candidate) error {
	return service.candidateRepository.UpdateCandidate(ctx, id, candidate)
}

func (service *candidateService) ReadCandidate(ctx context.Context, id string) (model.Candidate, error) {
	return service.candidateRepository.ReadCandidate(ctx, id)
}

func (service *candidateService) FindAllCandidates(ctx context.Context) ([]model.Candidate, error) {
	return service.candidateRepository.FindAllCandidates(ctx)
}

func (service *candidateService) FindCandidateByEmail(ctx context.Context, email string) (model.Candidate, error) {
	return service.candidateRepository.FindCandidateByEmail(ctx, email)
}

func (service *candidateService) FindAssigneesCandidates(ctx context.Context, id string) ([]model.Candidate, error) {
	// Check assignee exists with given id, return error if does not exist.
	a, _ := service.assigneeRepository.ReadAssignee(ctx, id)
	if a == (model.Assignee{}) {
		log.Println(model.ErrAssigneeDoesNotExist)
		return nil, model.ErrAssigneeDoesNotExist
	}

	return service.candidateRepository.FindAssigneesCandidates(ctx, id)
}

func (service *candidateService) DeleteCandidate(ctx context.Context, id string) error {
	// Check candidate exists with given id, return error if does not exist.
	c, _ := service.candidateRepository.ReadCandidate(ctx, id)
	if c == (model.Candidate{}) {
		log.Println(model.ErrCandidateDoesNotExist)
		return model.ErrCandidateDoesNotExist
	}

	return service.candidateRepository.DeleteCandidate(ctx, id)
}

func (service *candidateService) DenyCandidate(ctx context.Context, id string) error {
	// Check candidate exists with given id, return error if does not exist.
	c, _ := service.candidateRepository.ReadCandidate(ctx, id)
	if c == (model.Candidate{}) {
		log.Println(model.ErrCandidateDoesNotExist)
		return model.ErrCandidateDoesNotExist
	}

	c.Status = model.Denied

	return service.UpdateCandidate(ctx, id, c)
}

func (service *candidateService) AcceptCandidate(ctx context.Context, id string) error {
	// Check candidate exists with given id, return error if does not exist.
	c, _ := service.candidateRepository.ReadCandidate(ctx, id)
	if c == (model.Candidate{}) {
		log.Println(model.ErrCandidateDoesNotExist)
		return model.ErrCandidateDoesNotExist
	}

	// Candidates cannot be accepted before the completion of 4 meetings
	if c.MeetingCount < 4 {
		log.Println(model.ErrMeetingCountNotEnough)
		return model.ErrMeetingCountNotEnough
	}

	c.Status = model.Accepted

	return service.UpdateCandidate(ctx, id, c)
}

func (service *candidateService) ArrangeMeeting(ctx context.Context, id string, nextMeetingTime *time.Time) error {
	// Check candidate exists with given id, return error if does not exist.
	c, _ := service.candidateRepository.ReadCandidate(ctx, id)
	if c == (model.Candidate{}) {
		log.Println(model.ErrCandidateDoesNotExist)
		return model.ErrCandidateDoesNotExist
	}

	// Get assignees by department
	// Each time a random assignee is chosen by department
	a, _ := service.assigneeRepository.FindOneAssigneeByDepartment(ctx, c.Department)
	if a == (model.Assignee{}) {
		log.Println(model.ErrAssigneeDoesNotExist)
		return model.ErrAssigneeDoesNotExist
	}

	c.NextMeeting = nextMeetingTime

	// if this will not be the last meeting of the candidate
	// choose an assignee from the related department
	if c.MeetingCount >= 0 && c.MeetingCount < 3 {
		// Assign an assignee to the candidate according to the meeting count
		c.Assignee = a.ID

	} else if c.MeetingCount == 3 {
		// if current completed meeting number for the assignee is 3
		// which means the next meeting is the last (4th) one,
		// arrange next meeting with the CEO

		// Get the CEO
		ceo, _ := service.assigneeRepository.FindOneAssigneeByDepartment(ctx, model.CEO)
		if ceo == (model.Assignee{}) {
			log.Println(model.ErrAssigneeDoesNotExist)
			return model.ErrAssigneeDoesNotExist
		}
		// Assign the CEO to the candidate according in the last meeting
		c.Assignee = ceo.ID
	}

	return service.UpdateCandidate(ctx, id, c)
}

func (service *candidateService) CompleteMeeting(ctx context.Context, id string) error {
	// Check candidate exists with given id, return error if does not exist.
	c, _ := service.candidateRepository.ReadCandidate(ctx, id)
	if c == (model.Candidate{}) {
		log.Println(model.ErrCandidateDoesNotExist)
		return model.ErrCandidateDoesNotExist
	}

	// if the next meeting is null, it means current candidate does not have
	// any arranged meetings. then return an error accordingly.
	if c.NextMeeting == nil {
		log.Println(model.ErrArrangedMeetingDoesNotExist)
		return model.ErrArrangedMeetingDoesNotExist
	}

	// set next meeting to nil and update meeting count by one.
	c.NextMeeting = nil
	if c.MeetingCount < 4 {
		c.MeetingCount += 1
		c.Status = model.InProgress
	}

	return service.UpdateCandidate(ctx, id, c)
}
