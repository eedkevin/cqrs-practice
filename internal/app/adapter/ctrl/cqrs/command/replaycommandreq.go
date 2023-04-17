package command

import "time"

type ReplayCommandReq struct {
	RequestID     string `json:"request_id"`
	Command       string `json:"command"`
	Domain        string `json:"domain"`
	EventVelocity string `json:"event_velocity"`
	Payload       struct {
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
		ENV       string    `json:"env"`
	} `json:"payload"`
}
