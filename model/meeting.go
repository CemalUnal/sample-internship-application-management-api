package model

import (
	"time"
)

// Meeting model is used to exchange meeting metadata while arranging and completing meetings
// It is not persisted in the DB
type Meeting struct {
	CandidateID 	string 		`json:"candidate_id" validate:"required"`
	NextMeetingTime *time.Time	`json:"next_meeting_time" validate:"required"`
}
