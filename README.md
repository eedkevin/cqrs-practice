# CQRS Practise
A simplified practise for CQRS pattern and DDD in Golang

This repo uses a boilerplate from https://github.com/eedkevin/ddd-boilerplate-go that heavily inspired by https://github.com/resotto/goilerplate

## Folder Structure
```
.
├── cmd
│   └── app
│       ├── oddsworker  # oddsworker app entry
│       └── replayservice  # replayservice app entry
├── data  # data collection for local dev
├── design  # design docs
│   └── diagrams
├── internal  # internal packages
│   ├── app  # everything about the app
│   │   ├── adapter  # adapter layer
│   │   │   ├── ctrl  # controller impl
│   │   │   │   ├── cqrs  # cqrs controller impl
│   │   │   │   │   ├── command
│   │   │   │   │   └── query
│   │   │   │   └── healthz  # healthz controller impl
│   │   │   ├── repository  # repository impl
│   │   │   └── service  # service impl
│   │   ├── application  # application layer
│   │   │   ├── service  # service interface
│   │   │   └── usecase  # usecase impl, aka the core biz logics
│   │   ├── domain  # domain layer
│   │   │   ├── cqrs  # cqrs domain
│   │   │   │   └── model  # cqrs domain model
│   │   │   ├── event  # event domain
│   │   │   │   ├── model  # event domain model
│   │   │   │   └── repo  # event repo interface
│   │   │   └── odds  # odds domain
│   │   │       ├── model  # odds domain model
│   │   │       └── repo  # odds repo interface
│   │   └── infrastructure  # infrastructure layer
│   │       ├── mysql  # mysql
│   │       │   └── model  # mysql data model
│   │       ├── nats  # nats
│   │       └── sqlite3  # sqlite3
│   │           └── model  # sqlite3 data model
│   ├── cfg  # configs
│   └── util  # utils
└── testdata  # testing setup and mockups
```

## Quick Start
### Lift up the stack via docker compose
```sh
$ make compose-up
```

### Local Development

```sh
$ make install
$ make setup
$ make infra-up
$ make dev-replayservice
$ make dev-oddsworker
```

Note: hot reload is enabled via `make dev-replayservice` and `make dev-oddsworker`. any code changes will be triggering recomple. happy coding :)

