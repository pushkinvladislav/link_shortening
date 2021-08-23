.PHONY: all

include .env
export

PROJ_PATH := ${CURDIR}
DOCKER_PATH := ${PROJ_PATH}/docker

APP=link_shortening
MIGRATION_TOOL=goose
MIGRATIONS_DIR=./db/migrations

BASIC_IMAGE=dep
IMAGE_POSTFIX=-image

build:
	GOOS=linux GOARCH=arm go build -o .bin/client cmd/client/main.go
	GOOS=linux GOARCH=arm go build -o .bin/server cmd/server/main.go
	chmod ugo+x .bin/client
	chmod ugo+x .bin/server

build-docker:
	sudo rm -rf .database/
	docker build -t ${BASIC_IMAGE} -f ${DOCKER_PATH}/builder.Dockerfile.dev .
	docker build -t client${IMAGE_POSTFIX} -f ${DOCKER_PATH}/client.Dockerfile.dev .
	docker build -t server${IMAGE_POSTFIX} -f ${DOCKER_PATH}/server.Dockerfile.dev .

app-setup-and-up: build-docker app-up db-migrate-up 

app-up: build
	docker-compose up -d
	
all: app-setup-and-up

db-bash:
	docker-compose run --rm --no-deps --name link_shortening-db db ash

goose-init:
	GOOS=linux GOARCH=arm go build -o .bin/goose cmd/${MIGRATION_TOOL}/main.go
	chmod ugo+x .bin/${MIGRATION_TOOL}

db-up:
	docker-compose run --rm --no-deps --name link_shortening-db db ash

db-migration-create: goose-init
		goose -dir=${MIGRATIONS_DIR} create ${name} sql

db-migrate-status: goose-init
	docker-compose run --rm server .bin/goose -dir ${MIGRATIONS_DIR} postgres \
		"user=${POSTGRES_USER} host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${POSTGRES_SSL}" status

db-migrate-up: goose-init
	docker-compose run --rm server .bin/goose -dir ${MIGRATIONS_DIR} postgres \
        "user=${POSTGRES_USER} host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${POSTGRES_SSL}" up

db-migrate-down: goose-init
	docker-compose run --rm server .bin/goose -dir ${MIGRATIONS_DIR} postgres \
        "user=${POSTGRES_USER} host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${POSTGRES_SSL}" down

test:
	gotest -v ./...

proto: 
	# rm api/shorter/*.go
	protoc -I api/proto --go_out=plugins=grpc:api/shorter api/proto/shorter.proto

run: 
	go run cmd/client/main.go
