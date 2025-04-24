# syntax=docker/dockerfile:1

FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.mod ./
COPY dbconfig.yml ./
COPY ddl /app/ddl

RUN go mod tidy && \
	go install github.com/rubenv/sql-migrate/...@latest

CMD ["sql-migrate", "up", "-env=deploy" ]
