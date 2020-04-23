package api

import (
	"bytes"
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/cemalunal/sample-internship-management-api/model/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func assertHelper (t *testing.T, r *mux.Router, method string, path string, body []byte, expectedCode int) {
	request, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, expectedCode, response.Code)
}

func sendPostAndExpectOk(t *testing.T, r *mux.Router, path string, body []byte) {
	assertHelper(t, r, "POST", path, body, 200)
}

func sendPostAndExpectCreated(t *testing.T, r *mux.Router, path string, body []byte) {
	assertHelper(t, r, "POST", path, body, 201)
}

func sendPostAndExpectBadRequest(t *testing.T, r *mux.Router, path string, body []byte) {
	assertHelper(t, r, "POST", path, body, 400)
}

func sendGetAndExpectOk(t *testing.T, r *mux.Router, path string) {
	assertHelper(t, r, "GET", path, nil, 200)
}

func sendGetAndExpectBadRequest(t *testing.T, r *mux.Router, path string) {
	assertHelper(t, r, "GET", path, nil, 400)
}

func sendDeleteAndExpectOk(t *testing.T, r *mux.Router, path string) {
	assertHelper(t, r, "DELETE", path, nil, 200)
}

func sendDeleteAndExpectBadRequest(t *testing.T, r *mux.Router, path string) {
	assertHelper(t, r, "DELETE", path, nil, 400)
}

func sendPatchAndExpectOk(t *testing.T, r *mux.Router, path string) {
	assertHelper(t, r, "PATCH", path, nil, 200)
}

func sendPatchAndExpectBadRequest(t *testing.T, r *mux.Router, path string) {
	assertHelper(t, r, "PATCH", path, nil, 400)
}

func createCandidateSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates", mockApi.CreateCandidate).Methods(http.MethodPost)
	return router
}

func createCandidateAlreadyExistsRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceAlreadyExistsErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates", mockApi.CreateCandidate).Methods(http.MethodPost)
	return router
}

func readCandidateSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/{id}", mockApi.ReadCandidate).Methods(http.MethodGet)
	return router
}

func deleteCandidateSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/{id}", mockApi.DeleteCandidate).Methods(http.MethodDelete)
	return router
}

func deleteCandidateDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceDoesNotExistErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/{id}", mockApi.DeleteCandidate).Methods(http.MethodDelete)
	return router
}

func denyCandidateSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/deny/{id}", mockApi.DenyCandidate).Methods(http.MethodPatch)
	return router
}

func denyCandidateDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceDoesNotExistErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/deny/{id}", mockApi.DenyCandidate).Methods(http.MethodPatch)
	return router
}

func acceptCandidateSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/accept/{id}", mockApi.AcceptCandidate).Methods(http.MethodPatch)
	return router
}

func acceptCandidateDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceDoesNotExistErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/accept/{id}", mockApi.AcceptCandidate).Methods(http.MethodPatch)
	return router
}

func acceptCandidateNotEnoughMeetingRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceNotEnoughMeetingErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/accept/{id}", mockApi.AcceptCandidate).Methods(http.MethodPatch)
	return router
}

func findAllCandidatesSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates", mockApi.FindAllCandidates).Methods(http.MethodGet)
	return router
}

func findAssigneesCandidatesSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/assigneeId/{assigneeId}", mockApi.FindAssigneesCandidates).Methods(http.MethodGet)
	return router
}

func findAssigneesCandidatesAssigneeDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceDoesNotExistErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/candidates/assigneeId/{assigneeId}", mockApi.FindAssigneesCandidates).Methods(http.MethodGet)
	return router
}

func createAssigneeSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/assignees", mockApi.CreateAssignee).Methods(http.MethodPost)
	return router
}

func findAllAssigneesSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/assignees", mockApi.FindAllAssignees).Methods(http.MethodGet)
	return router
}

func findAssigneeIdByNameSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/assignees/name/{name}", mockApi.FindAssigneeIDByName).Methods(http.MethodGet)
	return router
}

func findAssigneeIdByNameAssigneeDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeServiceDoesNotExistErr(),
	}
	router.HandleFunc("/assignees/name/{name}", mockApi.FindAssigneeIDByName).Methods(http.MethodGet)
	return router
}

func findAllAssigneesByDepartmentSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/assignees/department/{department}",
		mockApi.FindAllAssigneesByDepartment).Methods(http.MethodGet)
	return router
}

func arrangeMeetingSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/meetings/arrange", mockApi.ArrangeMeeting).Methods(http.MethodPost)
	return router
}

func arrangeMeetingCandidateDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceDoesNotExistErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/meetings/arrange", mockApi.ArrangeMeeting).Methods(http.MethodPost)
	return router
}

func completeMeetingSuccessRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateService(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/meetings/complete/{candidateId}", mockApi.CompleteMeeting).Methods(http.MethodPost)
	return router
}

func completeMeetingCandidateDoesNotExistRouter() *mux.Router {
	router := mux.NewRouter()
	mockApi := api{
		CandidateService: mockCandidateServiceDoesNotExistErr(),
		AssigneeService:  mockAssigneeService(),
	}
	router.HandleFunc("/meetings/complete/{candidateId}", mockApi.CompleteMeeting).Methods(http.MethodPost)
	return router
}

func mockCandidateService() *mocks.CandidateService{
	candidate := mockCandidateModel()
	mockCandidateService := new(mocks.CandidateService)
	mockCandidateService.On("ReadCandidate", mock.Anything, mock.AnythingOfType("string")).Return(candidate, nil).Once()
	mockCandidateService.On("FindAllCandidates", mock.Anything).Return(mockCandidateArray(), nil).Once()
	mockCandidateService.On("CreateCandidate", mock.Anything, candidate).Return(candidate, nil).Once()
	mockCandidateService.On("DeleteCandidate", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
	mockCandidateService.On("DenyCandidate", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
	mockCandidateService.On("AcceptCandidate", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
	mockCandidateService.On("FindAssigneesCandidates", mock.Anything, mock.AnythingOfType("string")).
		Return(mockCandidateArray(), nil).Once()
	mockCandidateService.On("ArrangeMeeting", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()
	mockCandidateService.On("CompleteMeeting", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

	return mockCandidateService
}

func mockCandidateServiceDoesNotExistErr() *mocks.CandidateService{
	mockCandidateService := new(mocks.CandidateService)
	mockCandidateService.On("DeleteCandidate", mock.Anything, mock.AnythingOfType("string")).
		Return(model.ErrCandidateDoesNotExist).Once()
	mockCandidateService.On("DenyCandidate", mock.Anything, mock.AnythingOfType("string")).
		Return(model.ErrCandidateDoesNotExist).Once()
	mockCandidateService.On("AcceptCandidate", mock.Anything, mock.AnythingOfType("string")).
		Return(model.ErrCandidateDoesNotExist).Once()
	mockCandidateService.On("FindAssigneesCandidates", mock.Anything, mock.AnythingOfType("string")).
		Return([]model.Candidate{}, model.ErrAssigneeDoesNotExist).Once()
	mockCandidateService.On("ArrangeMeeting", mock.Anything, mock.AnythingOfType("string"), mock.Anything).
		Return(model.ErrCandidateDoesNotExist).Once()
	mockCandidateService.On("CompleteMeeting", mock.Anything, mock.AnythingOfType("string")).
		Return(model.ErrCandidateDoesNotExist).Once()

	return mockCandidateService
}

func mockCandidateServiceNotEnoughMeetingErr() *mocks.CandidateService{
	mockCandidateService := new(mocks.CandidateService)

	mockCandidateService.On("AcceptCandidate", mock.Anything, mock.AnythingOfType("string")).
		Return(model.ErrMeetingCountNotEnough).Once()

	return mockCandidateService
}
func mockCandidateServiceAlreadyExistsErr() *mocks.CandidateService {
	candidate := mockCandidateModel()
	mockCandidateService := new(mocks.CandidateService)
	mockCandidateService.On("CreateCandidate", mock.Anything, candidate).Return(model.Candidate{}, model.ErrCandidateAlreadyExists).Once()

	return mockCandidateService
}

func mockAssigneeService() *mocks.AssigneeService {
	assignee := mockAssigneeModel()
	mockAssigneeService := new(mocks.AssigneeService)
	mockAssigneeService.On("CreateAssignee", mock.Anything, assignee).Return(assignee, nil).Once()
	mockAssigneeService.On("FindAllAssignees", mock.Anything).Return(mockAssigneeArray(), nil).Once()
	mockAssigneeService.On("FindAllAssigneesByDepartment", mock.Anything, mock.AnythingOfType("string")).
		Return(mockAssigneeArray(), nil).Once()
	mockAssigneeService.On("FindAssigneeIDByName", mock.Anything, mock.AnythingOfType("string")).
		Return(assignee.ID, nil).Once()

	return mockAssigneeService
}

func mockAssigneeServiceDoesNotExistErr() *mocks.AssigneeService {
	mockAssigneeService := new(mocks.AssigneeService)
	mockAssigneeService.On("FindAssigneeIDByName", mock.Anything, mock.AnythingOfType("string")).
		Return("", model.ErrAssigneeDoesNotExist).Once()

	return mockAssigneeService
}

func mockCandidateModel() model.Candidate {
	return model.Candidate{
		ID: "asd",
		FirstName: "FN",
		LastName: "LN",
		Email: "e@e.com",
		Department: model.Development,
		University: "HU",
		Experience: false,
	}
}

func mockAssigneeModel() model.Assignee {
	return model.Assignee{
		ID: "asd",
		Name: "A1",
		Department: model.Development,
	}
}

func mockMeetingModel() model.Meeting {
	nextMeetingTime, _ := time.Parse(time.RFC3339, "2020-05-03T13:40:00.000+00:00")
	return model.Meeting{
		CandidateID: "asd",
		NextMeetingTime: &nextMeetingTime,
	}
}

func mockCandidateArray() []model.Candidate {
	candidateArray := make([]model.Candidate, 5)
	for i := 0; i < len(candidateArray); i++ {
		candidateArray[i] = mockCandidateModel()
	}

	return candidateArray
}

func mockAssigneeArray() []model.Assignee {
	assigneeArray := make([]model.Assignee, 5)
	for i := 0; i < len(assigneeArray); i++ {
		assigneeArray[i] = mockAssigneeModel()
	}

	return assigneeArray
}
