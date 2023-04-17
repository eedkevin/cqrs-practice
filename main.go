package main

import (
	"cqrs-practise/cmd/app/oddsworker"
	"cqrs-practise/cmd/app/replayservice"
	"cqrs-practise/internal/cfg"
	"flag"
	"log"
)

var (
	app = flag.String("app", "", "app name")
)

func main() {
	flag.Parse()
	cfg.Init()

	switch *app {
	case "replayservice":
		log.Println("Bootstrap replayservice")
		replayservice.Bootstrap()
	case "oddsworker":
		log.Println("Bootstrap oddsworker")
		oddsworker.Bootstrap()
	default:
		log.Fatalf("Unsupported app: %s", *app)
	}
}
