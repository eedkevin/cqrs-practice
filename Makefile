.PHONY: install
install:
	go install github.com/cosmtrek/air@latest

.PHONY: setup
setup:
	cp .env.template .env

.PHONY: test
test:
	go test -v ./...

.PHONY: up-oddsworker
up:
	go run main.go -app=oddsworker

.PHONY: up-replayservice
up-replayservice:
	go run main.go -app=replayservice

.PHONY: dev-oddsworker
dev-oddsworker:
	air -c .oddsworker.air.toml

.PHONY: dev-replayservice
dev-replayservice:
	air -c .replayservice.air.toml

.PHONY: infra-up
infra-up:
	docker compose up nats -d

.PHONY: infra-down
infra-down:
	docker compose down nats

.PHONY: stack-up
stack-up:
	docker compose up --build

.PHONY: stack-down
stack-down:
	docker compose down