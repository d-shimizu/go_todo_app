.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest

build: ## Build docker image to deploy
	docker build -t localhost/go_todo_app:$(DOCKER_TAG) \
		--target deploy ./

build-local: ## Build docker image to local deployment
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up --build

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check containers status
	docker compose ps

test: ## Excute tests
	go test -race -shuffle=on ./...

dry-migrate: ## Try Migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo --dry-run < ./_tools/mysql/schema.sql

migrate: ## Exacute Migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo < ./_tools/mysql/schema.sql

help: ## show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
