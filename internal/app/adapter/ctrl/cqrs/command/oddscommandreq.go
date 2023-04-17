package command

type OddsCommandReq struct {
	RequestID     string `json:"request_id"`
	Command       string `json:"command"`
	Domain        string `json:"domain"`
	EventVelocity string `json:"event_velocity"`
	Payload       struct {
		GameUUID  string `json:"game_uuid"`
		MoneyLine struct {
			Home float64 `json:"home"`
			Away float64 `json:"away"`
			Draw float64 `json:"draw"`
		} `json:"moneyline"`
	} `json:"payload"`
}
