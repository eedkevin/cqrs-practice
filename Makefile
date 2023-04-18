.PHONY: install
install:
	go install github.com/cosmtrek/air@latest

.PHONY: setup
setup:
	cp .env.template .env && sed -i "" -e 's/{{sqlite-db}}/.\/data\/data.db/g' .env && sed -i "" -e 's/{{nats-url}}/nats:\/\/127.0.0.1:4222/g' .env
	cp .env.template .env.docker && sed -i "" -e 's/{{sqlite-db}}/.\/data.db/g' .env.docker && sed -i "" -e 's/{{nats-url}}/nats:\/\/nats-server/g' .env.docker

.PHONY: test
test:
	go test -v ./...

.PHONY: up-oddsworker
up-oddsworker:
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
	docker compose up nats-server -d

.PHONY: infra-down
infra-down:
	docker compose down nats-server

.PHONY: stack-up
stack-up:
	docker compose up --build

.PHONY: stack-down
stack-down:
	docker compose down