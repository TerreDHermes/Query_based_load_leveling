.PHONY: all build run start

LOCAL_PORT:=5436
DOCKER_PORT:=5432
PASSWORD:=qwerty
USERNAME:=postgres
HOST:=localhost
SSLMODE:=disable
DATABASE:=postgres

all: build run start

build:
	docker build -t postgres .

run:
	docker run --name postgres -p $(LOCAL_PORT):$(DOCKER_PORT) -e POSTGRES_PASSWORD=$(PASSWORD) -d postgres

start:
	go run cmd/backend/main.go
