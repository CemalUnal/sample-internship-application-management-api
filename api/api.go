package api

import (
	"github.com/cemalunal/sample-internship-management-api/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type api struct {
	AssigneeService model.AssigneeService
	CandidateService model.CandidateService
}

func Api(router *mux.Router, assigneeService model.AssigneeService, candidateService model.CandidateService) *mux.Router {
	_api := &api{
		AssigneeService: assigneeService,
		CandidateService: candidateService,
	}

	router.HandleFunc("/candidates", _api.CreateCandidate).Methods(http.MethodPost)
	router.HandleFunc("/candidates", _api.FindAllCandidates).Methods(http.MethodGet)
	router.HandleFunc("/candidates/{id}", _api.ReadCandidate).Methods(http.MethodGet)
	router.HandleFunc("/candidates/{id}", _api.DeleteCandidate).Methods(http.MethodDelete)
	router.HandleFunc("/candidates/deny/{id}", _api.DenyCandidate).Methods(http.MethodPatch)
	router.HandleFunc("/candidates/accept/{id}", _api.AcceptCandidate).Methods(http.MethodPatch)
	router.HandleFunc("/candidates/assigneeId/{assigneeId}", _api.FindAssigneesCandidates).Methods(http.MethodGet)
	router.HandleFunc("/assignees", _api.CreateAssignee).Methods(http.MethodPost)
	router.HandleFunc("/assignees", _api.FindAllAssignees).Methods(http.MethodGet)
	router.HandleFunc("/assignees/name/{name}", _api.FindAssigneeIDByName).Methods(http.MethodGet)
	router.HandleFunc("/assignees/department/{department}", _api.FindAllAssigneesByDepartment).Methods(http.MethodGet)
	router.HandleFunc("/meetings/arrange", _api.ArrangeMeeting).Methods(http.MethodPost)
	router.HandleFunc("/meetings/complete/{candidateId}", _api.CompleteMeeting).Methods(http.MethodPost)
	router.Use(RequestLogger)

	log.Fatalln(http.ListenAndServe(":8080", router))
	return router
}
