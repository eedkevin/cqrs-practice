package model

type OddsCommand struct {
	Command
	Payload struct {
		GameID    string
		MoneyLine struct {
			Home float64
			Away float64
			Draw float64
		}
	}
}
