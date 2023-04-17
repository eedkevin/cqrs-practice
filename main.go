package main

import (
	"cqrs-practise/cmd/app/oddsworker"
	"cqrs-practise/cmd/app/replayservice"
	"flag"
	"log"
)

var (
	app = flag.String("app", "", "app name")
)

func main() {
	flag.Parse()

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
