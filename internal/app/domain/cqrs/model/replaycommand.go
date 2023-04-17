package model

import "time"

type ReplayCommand struct {
	Command
	Payload struct {
		StartTime time.Time
		EndTime   time.Time
		ENV       string
	}
}
