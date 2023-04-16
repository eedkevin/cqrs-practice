package model

type Odds struct {
	UUID      string
	GameUUID  string
	MoneyLine struct {
		Home float64
		Away float64
		Draw float64
	}
}
