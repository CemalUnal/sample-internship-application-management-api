FROM golang:1.14.2-alpine AS builder

WORKDIR /go/src/sample-internship-management-api
COPY . .
RUN go build -o /sample-internship-management-api .

FROM alpine AS release

WORKDIR /sample-internship-management-api
COPY --from=builder /sample-internship-management-api ./server

EXPOSE 8080

ENTRYPOINT ["/sample-internship-management-api/server"]
