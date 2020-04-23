package main

import (
	"github.com/cemalunal/sample-internship-management-api/api"
	"github.com/cemalunal/sample-internship-management-api/db"
	_assigneeRepository "github.com/cemalunal/sample-internship-management-api/repository"
	_candidateRepository "github.com/cemalunal/sample-internship-management-api/repository"
	_assigneeService "github.com/cemalunal/sample-internship-management-api/service"
	_candidateService "github.com/cemalunal/sample-internship-management-api/service"
	"github.com/gorilla/mux"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mongodbUri := os.Getenv("MONGODB_URI")
	if mongodbUri == "" {
		log.Println("MONGODB_URI is not set, using default connection string mongodb://localhost:27017")
		mongodbUri = "mongodb://localhost:27017"
	}

	client := db.Connect(mongodbUri)
	database := client.Database("Company")
	assigneesCollection := database.Collection("Assignees")
	candidatesCollection := database.Collection("Candidates")

	assigneeRepository := _assigneeRepository.MongoDBAssigneeRepository(assigneesCollection)
	candidateRepository := _candidateRepository.MongoDBCandidateRepository(candidatesCollection)
	assigneeService := _assigneeService.AssigneeService(assigneeRepository)
	candidateService := _candidateService.CandidateService(candidateRepository, assigneeRepository)

	r := mux.NewRouter()
	api.Api(r, assigneeService, candidateService)
}
