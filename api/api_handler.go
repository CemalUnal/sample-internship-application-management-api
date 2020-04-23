package api

import (
	"encoding/json"
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

// CreateCandidate creates candidate by given request body
func (a *api) CreateCandidate(w http.ResponseWriter, req *http.Request) {
	// create candidate model from request body
	var candidate model.Candidate
	err := json.NewDecoder(req.Body).Decode(&candidate)
	if err != nil {
		a.ReturnBadRequest(w, err)
		return
	}

	// try to validate the fields of the candidate
	if ok, err := a.IsRequestValid(candidate); !ok {
		a.ReturnBadRequest(w, err)
		return
	}

	// Check given department is in the existing departments
	// and do not allow to create candidate if not
	departmentIsValid := a.CheckDepartmentExists(candidate.Department)
	if !departmentIsValid {
		a.ReturnBadRequest(w, model.ErrDepartmentDoesNotExist)
		return
	}

	createdCandidate, err := a.CandidateService.CreateCandidate(req.Context(), candidate)
	if err != nil {
		if err == model.ErrCandidateAlreadyExists {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnCreated(w, "Successfully created candidate", createdCandidate)
	log.Println("Successfully created candidate with id: ", createdCandidate.ID)
}

// FindAllCandidates finds all candidates that are available in the system
func (a *api) FindAllCandidates(w http.ResponseWriter, req *http.Request) {
	candidates, err := a.CandidateService.FindAllCandidates(req.Context())
	if err != nil {
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully fetched all candidates", candidates)
	log.Println("Successfully fetched all candidates.")
}

// ReadCandidate finds a candidate by given id
func (a *api) ReadCandidate(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	candidate, err := a.CandidateService.ReadCandidate(req.Context(), id)
	if err != nil {
		a.ReturnBadRequest(w, err)
		return
	}

	a.ReturnOk(w, "Successfully read candidate", candidate)
	log.Println("Successfully read candidate with id: ", id)
}

// DeleteCandidate deletes a candidate by given id
func (a *api) DeleteCandidate(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	err := a.CandidateService.DeleteCandidate(req.Context(), id)
	if err != nil {
		if err == model.ErrCandidateDoesNotExist {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully deleted candidate", id)
	log.Println("Successfully deleted candidate with id: ", id)
}

// DenyCandidate denies a candidate by given id
func (a *api) DenyCandidate(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	err := a.CandidateService.DenyCandidate(req.Context(), id)
	if err != nil {
		if err == model.ErrCandidateDoesNotExist {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully denied candidate", id)
	log.Println("Successfully denied candidate with id: ", id)
}

// AcceptCandidate accepts a candidate by given id
func (a *api) AcceptCandidate(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	err := a.CandidateService.AcceptCandidate(req.Context(), id)
	if err != nil {
		if err == model.ErrCandidateDoesNotExist || err == model.ErrMeetingCountNotEnough {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully accepted candidate", id)
	log.Println("Successfully accepted candidate with id: ", id)
}

// FindAssigneesCandidates finds the assignee's candidates by given assignee id
func (a *api) FindAssigneesCandidates(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	assigneeId := params["assigneeId"]

	candidates, err := a.CandidateService.FindAssigneesCandidates(req.Context(), assigneeId)
	if err != nil {
		if err == model.ErrAssigneeDoesNotExist {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully found candidates of assignee", candidates)
	log.Println("Successfully found candidates of assignee with id: ", assigneeId)
}

// FindAllAssignees finds all assignees that are available in the system
func (a *api) FindAllAssignees(w http.ResponseWriter, req *http.Request) {
	assignees, err := a.AssigneeService.FindAllAssignees(req.Context())
	if err != nil {
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully fetched all assignees", assignees)
	log.Println("Successfully fetched all assignees.")
}

// CreateAssignee creates assignee by given request body
func (a *api) CreateAssignee(w http.ResponseWriter, req *http.Request) {
	// create assignee model from request body
	var assignee model.Assignee
	err := json.NewDecoder(req.Body).Decode(&assignee)
	if err != nil {
		a.ReturnBadRequest(w, err)
		return
	}

	// try to validate the fields of the assignee
	if ok, err := a.IsRequestValid(assignee); !ok {
		a.ReturnBadRequest(w, err)
		return
	}

	// Check given department is in the existing departments
	// and do not allow to create assignee if not
	departmentIsValid := a.CheckDepartmentExists(assignee.Department)
	if !departmentIsValid {
		a.ReturnBadRequest(w, model.ErrDepartmentDoesNotExist)
		return
	}

	createdAssignee, err := a.AssigneeService.CreateAssignee(req.Context(), assignee)
	if err != nil {
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnCreated(w, "Successfully created assignee", createdAssignee)
	log.Println("Successfully created assignee with id: ", createdAssignee.ID)
}

// FindAssigneeIDByName finds assignee id by given assignee name
func (a *api) FindAssigneeIDByName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	assigneeName := params["name"]

	id := a.AssigneeService.FindAssigneeIDByName(req.Context(), assigneeName)
	if id == "" {
		a.ReturnBadRequest(w, model.ErrAssigneeDoesNotExist)
		return
	}

	a.ReturnOk(w, "Successfully found assignee id by name", id)
	log.Println("Successfully found assignee id by name: ", assigneeName)
}

// FindAllAssigneesByDepartment finds assignee id by given assignee name
func (a *api) FindAllAssigneesByDepartment(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	department := params["department"]

	// Check given department is in the existing departments
	// and do not allow to create candidate if not
	departmentIsValid := a.CheckDepartmentExists(department)
	if !departmentIsValid {
		a.ReturnBadRequest(w, model.ErrDepartmentDoesNotExist)
		return
	}

	assignees, err := a.AssigneeService.FindAllAssigneesByDepartment(req.Context(), department)
	if err != nil {
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully fetched all assignees by department", assignees)
	log.Println("Successfully fetched all assignees by department: ", department)
}

// ArrangeMeeting arranges a meeting with the given candidate on the given date
func (a *api) ArrangeMeeting(w http.ResponseWriter, req *http.Request) {
	// create meeting model from request body
	var meeting model.Meeting
	err := json.NewDecoder(req.Body).Decode(&meeting)
	if err != nil {
		a.ReturnBadRequest(w, err)
		return
	}
	// try to validate the fields of the meeting
	if ok, err := a.IsRequestValid(meeting); !ok {
		a.ReturnBadRequest(w, err)
		return
	}

	err = a.CandidateService.ArrangeMeeting(req.Context(), meeting.CandidateID, meeting.NextMeetingTime)
	if err != nil {
		if err == model.ErrCandidateDoesNotExist || err == model.ErrAssigneeDoesNotExist {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully arranged meeting with candidate on given date", meeting)
	log.Printf("Successfully arranged meeting with candidate with id: %s, on date %s\n",
		meeting.CandidateID, meeting.NextMeetingTime)
}

// CompleteMeeting completes a meeting of a candidate by given candidate id
func (a *api) CompleteMeeting(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	candidateId := params["candidateId"]

	err := a.CandidateService.CompleteMeeting(req.Context(), candidateId)
	if err != nil {
		if err == model.ErrCandidateDoesNotExist || err == model.ErrArrangedMeetingDoesNotExist {
			a.ReturnBadRequest(w, err)
			return
		}
		a.ReturnInternalServerError(w, err)
		return
	}

	a.ReturnOk(w, "Successfully completed meeting with candidate", candidateId)
	log.Println("Successfully completed meeting with candidate with id: ", candidateId)
}

// EncodeApiResponse is a helper function to create response body as json
func (a *api) EncodeApiResponse(w http.ResponseWriter, response model.ApiResponse) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
}

// ReturnInternalServerError is a helper function to return Internal Server Error response with given error message
func (a *api) ReturnInternalServerError(w http.ResponseWriter, err error) {
	var response model.ApiResponse
	log.Println(err)
	response, w = model.GetInternalServerErrorResponse(w, err.Error())

	a.EncodeApiResponse(w, response)
}

// ReturnBadRequest is a helper function to return Bad Request response with given error message
func (a *api) ReturnBadRequest(w http.ResponseWriter, err error) {
	var response model.ApiResponse
	log.Println(err)
	response, w = model.GetBadRequestResponse(w, err.Error())

	a.EncodeApiResponse(w, response)
}

// ReturnCreated is a helper function to return Created response with given response body
func (a *api) ReturnCreated(w http.ResponseWriter, message string, responseBody interface{}) {
	var response model.ApiResponse
	response, w = model.GetCreatedResponse(w, message, responseBody)

	a.EncodeApiResponse(w, response)
}

// ReturnOk is a helper function to return Ok response with given response body
func (a *api) ReturnOk(w http.ResponseWriter, message string, responseBody interface{}) {
	var response model.ApiResponse
	response, w = model.GetOkResponse(w, message, responseBody)

	a.EncodeApiResponse(w, response)
}

// IsRequestValid is a helper function to validate request body
func (a *api) IsRequestValid(body interface{}) (bool, error) {
	v := validator.New()
	err := v.Struct(body)
	if err != nil {
		return false, err
	}

	return true, nil
}

// CheckDepartmentExists is a helper function to check the given department exists in the system
func (a *api) CheckDepartmentExists(department string) bool {
	departments := model.GetDepartmentsAsArray()
	departmentIsValid := false
	for i := 0;  i < len(departments); i++ {
		if strings.Compare(departments[i], department) == 0 {
			departmentIsValid = true
		}
	}

	return departmentIsValid
}
