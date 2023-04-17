package command

type CommandReq struct {
	RequestID     string      `json:"request_id"`
	Command       string      `json:"command"`
	Domain        string      `json:"domain"`
	EventVelocity string      `json:"event_velocity"`
}
