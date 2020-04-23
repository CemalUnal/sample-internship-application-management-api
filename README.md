# **Sample Internship Management Api**

## Repository Structure

- [api](./api) layer contains the rest api endpoints and handlers that allows us to manage candidates, their applications, meetings and assignees

- [db](./db) layer connects to the MongoDB with given connection string.

- [repository](./repository) layer executes CRUD queries on the database only. It does not contain any business logic.

- [service](./service) layer contains the core business logic. It handles communication between api and repository layers.

- [model](./model) layer contains the models that we want to persist to the DB. For example : Assignee, Candidate. Also, service and repository interfaces defined in this layer.

- [model/mocks](./model/mocks) layer contains the mock implementation for the repository and service layers that are consumed by the tests.

- [main.go](./main.go) main.go

- [mongo-seed](./mongo-seed) mongo-seed contains the sample data (Data provided in the task sheet), and a script that allows us to import the sample data.

- [docker-compose.yml](./docker-compose.yml) contains the configurations that allows us to run the rest api along with single MongoDB instance with sample data using Docker.

## Data Model

- [Candidate](./model/candidate.go)
    - This model used to store and exchange candidate information. It is persisted in the DB in Candidates collection.
    - Also, the related service and repository interfaces declared in the same file along with this model.
- [Assignee](./model/assignee.go)
    - This model used to store and exchange assignee information. It is persisted in the DB in Assignees collection.
    - Also, the related service and repository interfaces declared in the same file along with this model.
- [Department](./model/department.go)
    - This model used to simulate enumeration for the Department info, and it is not persisted in the DB.
- [Status](./model/status.go)
    - This model used to simulate enumeration for the Status info, and it is not persisted in the DB.
- [Meeting](./model/meeting.go)
    - This model used to exchange meeting metadata while arranging and completing meetings, and it is not persisted in the DB.

## Running
### Quick Start with Docker Compose

### Prerequisites

- docker 19.03.2+
- docker-compose 1.23.2+

(If you have decided to run the application using Docker Compose, then you can skip [Alternatively Running with Docker Commands](#alternatively-running-with-docker-commands) part and continue with the [Usage](#usage) part.)

Execute the following command to start the application:

```bash
docker-compose up -d --build
```

Above command will build an image for the rest api and then run it along with standalone MongoDB instance with sample data.

Check everything is working properly:

```bash
docker ps

CONTAINER ID        IMAGE                                                               COMMAND                  CREATED             STATUS              PORTS                      NAMES
b0dc6e9cc6d4        sample-internship-management-api_sample-internship-management-api   "/sample-internship-…"   17 seconds ago      Up 16 seconds       0.0.0.0:8080->8080/tcp     sample-internship-management-api
4e0cb722d253        mongo:3.6.10                                                        "docker-entrypoint.s…"   17 seconds ago      Up 16 seconds       0.0.0.0:27017->27017/tcp   mongodb
```
Please note that each time when you create the MongoDB container, existing database will be dropped, and a new one will be created. If you do not want to drop the existing database, you can remove the `--drop` parameter from the [import.sh](./mongo-seed/import.sh) script.

If you want to stop and remove the running containers:

```bash
docker-compose down
```

### Alternatively Running with Docker Commands

### Prerequisites

- docker 19.03.2+

Create a network that is called dev-network:

```bash
docker network create dev-network
```

Create a volume for MongoDB and run its container.

```bash
docker volume create --name mongodb_data
docker run -p 27017:27017 -d --network=dev-network \
            --name mongodb \
            -v mongodb_data:/data/db \
            --restart=on-failure \
            mongo:3.6.10
```

Create sample db with `mongorestore`:

```bash
docker run -it --rm -v $(pwd)/mongo-seed:/mongo-seed \
  --network=dev-network mongo:3.6.10 /mongo-seed/import.sh
```

Build Docker image for the rest api:

```
docker build -t sample-internship-management-api .
```

```bash
docker run -p 8080:8080 -d --network=dev-network \
             -e MONGODB_URI=mongodb://mongodb:27017 \
             --name sample-internship-management-api \
             --restart=on-failure \
             sample-internship-management-api
```

Check everything is working properly:

```bash
docker ps

CONTAINER ID        IMAGE                              COMMAND                  CREATED             STATUS              PORTS                      NAMES
f02936f57f48        sample-internship-management-api   "/sample-internship-…"   7 seconds ago       Up 6 seconds        0.0.0.0:8080->8080/tcp     sample-internship-management-api
56f4a71e3e43        mongo:3.6.10                       "docker-entrypoint.s…"   43 seconds ago      Up 42 seconds       0.0.0.0:27017->27017/tcp   mongodb
```

If you want to stop and remove the running containers:

```bash
docker stop sample-internship-management-api
docker stop mongodb
docker rm sample-internship-management-api
docker rm mongodb
```

## Usage

### Required Functions on the Task Sheet

#### Create Candidate

You can create a candidate by posting a candidate model like the following:
```bash
curl -X POST \
  http://localhost:8080/candidates \
  -H 'content-type: application/json' \
  -d '{
    "first_name" : "Cemal",
    "last_name" : "Unal",
    "email" : "test@test.com",
    "department" : "Development",
    "university" : "Ankara",
    "experience" : false
   }'
```
`ApplicationDate`, `Status`, `MeetingCount`, `NextMeeting` will be set automatically after you create the candidate. `Assignee` field will be set after you have arranged a meeting with this candidate. `Email`, `Department` and `University` fields required.
Also, email format should be example@email.xyz. Otherwise, the api returns bad request and candidate will be not inserted to DB.

#### Read Candidate

You can read a candidate by using its id like the following:
```bash
curl -X GET http://localhost:8080/candidates/5ea980281dafc611002fbc41
```

#### Delete Candidate

You can delete a candidate by using its id like the following:
```bash
curl -X DELETE http://localhost:8080/candidates/5ea980281dafc611002fbc41
```

#### Arrange Meeting

You can arrange a meeting with a candidate by posting candidate id and next meeting time in request body:
```bash
curl -X POST \
  http://localhost:8080/meetings/arrange \
  -H 'accept: application/json' \
  -d '{
	"candidate_id": "5ea980281dafc611002fbc41",
	"next_meeting_time": "2020-05-03T13:40:00.000+00:00"
  }'
```
Both of the `candidate_id` and `next_meeting_time` are required in order to arrange the meeting. After arranging a meeting, a randomly chosen assignee will be assigned to the given candidate according to the department they have applied.

#### Complete Meeting

You can complete meeting by posting the candidate id like the following:
```bash
curl -X POST http://localhost:8080/meetings/complete/5ea980281dafc611002fbc41
```

#### Deny Candidate

You can deny a candidate by using its id like the following:
```bash
curl -X PATCH http://localhost:8080/candidates/deny/5ea980281dafc611002fbc41
```

#### Accept Candidate

You can accept a candidate by using its id like the following:
```bash
curl -X PATCH http://localhost:8080/candidates/accept/5ea980281dafc611002fbc41
```

#### Find Assignee ID by Name

You can find the assignee id by using its name like the following:
```bash
curl -X GET http://localhost:8080/assignees/name/Sercan
```
Please note that this endpoint is case-sensitive. It **will not produce** the same result with sercan as it produced with Sercan

### Bonus Functions

#### Find Assignee's Candidates

You can find the candidates of the assignee by using its id like the following:
```bash
curl -X GET http://localhost:8080/candidates/assigneeId/5ea980281dafc611002fbc41
```

#### Find All Candidates

You can find all candidates that are available in the system like the following:
```bash
curl -X GET http://localhost:8080/candidates
```

#### Create Assignee

You can create an assignee by posting an assignee model like the following:
```bash
curl -X POST \
  http://localhost:8080/assignees \
  -H 'content-type: application/json' \
  -d '{
    "name" : "cemal",
    "department" : "Development"
   }'
```
Both of the `name` and `department` are required in order to create an assignee.

#### Find All Assignees

You can find all assignees that are available in the system like the following:
```bash
curl -X GET http://localhost:8080/assignees
```

#### Find All Assignees by Department

You can find all assignees in a department by using department name like the following:
```bash
curl -X GET http://localhost:8080/assignees/department/Design
```
Please note that this endpoint is case-sensitive. It **will not produce** the same result with design as it produced with Design

## Development

### Prerequisites

- Go 1.14.2+
- MongoDB 3.6.10+

### Running MongoDB for Development

You can easily start a MongoDB instance using Docker easily:

```bash
docker volume create --name mongodb_data
docker run -p 27017:27017 -d \
            --name mongodb \
            -v mongodb_data:/data/db \
            --restart=on-failure \
            mongo:3.6.10
```

Then you can run this rest api by executing the following command. You can specify the connection string for the newly created MongoDB instance using `MONGODB_URI` environment variable:

```bash
MONGODB_URI=mongodb://localhost:27017 go run .
```

**Note:** Application will try to connect to the MongoDB using the `mongodb://localhost:27017` as connection string if you do not set the MONGODB_URI environment variable.

### Running All Tests

```bash
go test ./...
```

### Generating Test Coverage Results

```bash
go test ./... -coverprofile coverage.out
?       github.com/cemalunal/sample-internship-management-api  [no test files]
ok      github.com/cemalunal/sample-internship-management-api/api      0.342s  coverage: 82.6% of statements
?       github.com/cemalunal/sample-internship-management-api/db       [no test files]
?       github.com/cemalunal/sample-internship-management-api/model    [no test files]
?       github.com/cemalunal/sample-internship-management-api/model/mocks      [no test files]
?       github.com/cemalunal/sample-internship-management-api/repository       [no test files]
ok      github.com/cemalunal/sample-internship-management-api/service  0.260s  coverage: 61.0% of statements
```
