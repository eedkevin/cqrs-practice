package main

import (
	"cqrs-practise/cmd/app/replayservice"
	"cqrs-practise/internal/cfg"
)

func main() {
	cfg.Init()
	replayservice.Bootstrap()
}
