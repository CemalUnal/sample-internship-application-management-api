package api

import (
	"encoding/json"
	"testing"
)

func TestApi_CreateCandidate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := createCandidateSuccessRouter()
		candidate := mockCandidateModel()
		jsonCandidate, _ := json.Marshal(candidate)
		sendPostAndExpectCreated(t, router, "/candidates", jsonCandidate)
	})

	t.Run("department-does-not-exists", func(t *testing.T) {
		router := createCandidateSuccessRouter()
		candidate := mockCandidateModel()
		candidate.Department = "test"
		jsonCandidate, _ := json.Marshal(candidate)
		sendPostAndExpectBadRequest(t, router, "/candidates", jsonCandidate)
	})

	t.Run("email-is-invalid", func(t *testing.T) {
		router := createCandidateSuccessRouter()
		candidate := mockCandidateModel()
		candidate.Email = "invalidEmail"
		jsonCandidate, _ := json.Marshal(candidate)
		sendPostAndExpectBadRequest(t, router, "/candidates", jsonCandidate)
	})

	t.Run("candidate-already-exists", func(t *testing.T) {
		router := createCandidateAlreadyExistsRouter()
		candidate := mockCandidateModel()
		jsonCandidate, _ := json.Marshal(candidate)
		sendPostAndExpectBadRequest(t, router, "/candidates", jsonCandidate)
	})
}

func TestApi_FindAllCandidates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := findAllCandidatesSuccessRouter()
		sendGetAndExpectOk(t, router, "/candidates")
	})
}

func TestApi_ReadCandidate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := readCandidateSuccessRouter()
		sendGetAndExpectOk(t, router, "/candidates/abcd")
	})
}

func TestApi_DeleteCandidate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := deleteCandidateSuccessRouter()
		sendDeleteAndExpectOk(t, router, "/candidates/abcd")
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		router := deleteCandidateDoesNotExistRouter()
		sendDeleteAndExpectBadRequest(t, router, "/candidates/abcd")
	})
}

func TestApi_DenyCandidate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := denyCandidateSuccessRouter()
		sendPatchAndExpectOk(t, router, "/candidates/deny/abcd")
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		router := denyCandidateDoesNotExistRouter()
		sendPatchAndExpectBadRequest(t, router, "/candidates/deny/abcd")
	})
}

func TestApi_AcceptCandidate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := acceptCandidateSuccessRouter()
		sendPatchAndExpectOk(t, router, "/candidates/accept/abcd")
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		router := acceptCandidateDoesNotExistRouter()
		sendPatchAndExpectBadRequest(t, router, "/candidates/accept/abcd")
	})

	t.Run("meeting-count-not-enough", func(t *testing.T) {
		router := acceptCandidateNotEnoughMeetingRouter()
		sendPatchAndExpectBadRequest(t, router, "/candidates/accept/abcd")
	})
}

func TestApi_FindAssigneesCandidates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := findAssigneesCandidatesSuccessRouter()
		sendGetAndExpectOk(t, router, "/candidates/assigneeId/abcd")
	})

	t.Run("assignee-does-not-exist", func(t *testing.T) {
		router := findAssigneesCandidatesAssigneeDoesNotExistRouter()
		sendGetAndExpectBadRequest(t, router, "/candidates/assigneeId/abcd")
	})
}

func TestApi_CreateAssignee(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := createAssigneeSuccessRouter()
		assignee := mockAssigneeModel()
		jsonAssignee, _ := json.Marshal(assignee)
		sendPostAndExpectCreated(t, router, "/assignees", jsonAssignee)
	})

	t.Run("department-does-not-exists", func(t *testing.T) {
		router := createAssigneeSuccessRouter()
		assignee := mockAssigneeModel()
		assignee.Department = "test"

		jsonAssignee, _ := json.Marshal(assignee)
		sendPostAndExpectBadRequest(t, router, "/assignees", jsonAssignee)
	})
}

func TestApi_FindAllAssignees(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := findAllAssigneesSuccessRouter()
		sendGetAndExpectOk(t, router, "/assignees")
	})
}

func TestApi_FindAssigneeIDByName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := findAssigneeIdByNameSuccessRouter()
		sendGetAndExpectOk(t, router, "/assignees/name/test")
	})

	t.Run("assignee-does-not-exist", func(t *testing.T) {
		router := findAssigneeIdByNameAssigneeDoesNotExistRouter()
		sendGetAndExpectBadRequest(t, router, "/assignees/name/test")
	})
}

func TestApi_FindAllAssigneesByDepartment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := findAllAssigneesByDepartmentSuccessRouter()
		sendGetAndExpectOk(t, router, "/assignees/department/Development")
	})
}

func TestApi_ArrangeMeeting(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := arrangeMeetingSuccessRouter()
		meeting := mockMeetingModel()
		jsonMeeting, _ := json.Marshal(meeting)
		sendPostAndExpectOk(t, router, "/meetings/arrange", jsonMeeting)
	})

	t.Run("field-required", func(t *testing.T) {
		router := arrangeMeetingSuccessRouter()
		meeting := mockMeetingModel()
		meeting.CandidateID = ""
		jsonMeeting, _ := json.Marshal(meeting)
		sendPostAndExpectBadRequest(t, router, "/meetings/arrange", jsonMeeting)
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		router := arrangeMeetingCandidateDoesNotExistRouter()
		meeting := mockMeetingModel()
		jsonMeeting, _ := json.Marshal(meeting)
		sendPostAndExpectBadRequest(t, router, "/meetings/arrange", jsonMeeting)
	})
}

func TestApi_CompleteMeeting(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := completeMeetingSuccessRouter()
		sendPostAndExpectOk(t, router, "/meetings/complete/qwe123", nil)
	})

	t.Run("candidate-does-not-exist", func(t *testing.T) {
		router := completeMeetingCandidateDoesNotExistRouter()
		sendPostAndExpectBadRequest(t, router, "/meetings/complete/qwe123", nil)
	})
}
