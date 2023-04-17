.PHONY: install
install:
	go install github.com/cosmtrek/air@latest

.PHONY: test
test:
	go test -v ./...

.PHONY: up
up:
	go run main.go

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
