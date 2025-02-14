package models

type TodayFixture struct {
	Leagues []LeagueFixture `json:"league"`
}

type LeagueFixture struct {
	League   string    `json:"league"`
	Fixtures []Fixture `json:"fixture"`
}

type Fixture struct {
	Referee   string   `json:"referee"`
	City      string   `json:"city"`
	Stadium   string   `json:"stadium"`
	HomeTeam  string   `json:"home_team"`
	AwayTeam  string   `json:"away_team"`
	LastGame  LastGame `json:"last_game"`
	HomeStats Stats    `json:"home_stats"`
	AwayStats Stats    `json:"away_stats"`
}

type LastGame struct {
	Date          string `json:"date"`
	City          string `json:"city"`
	Venue         string `json:"venue"`
	HomeTeam      string `json:"home"`
	AwayTeam      string `json:"away"`
	TeamWinner    string `json:"winner"`
	FinalScore    string `json:"final_score"`
	HalftimeScore string `json:"halftime_score"`
}

type Stats struct {
	WinsThisSeason     int    `json:"wins_this_season"`
	DrawsThisSeason    int    `json:"draws_this_season"`
	LosesThisSeason    int    `json:"loses_this_season"`
	TotalGoals         int    `json:"total_goals"`
	TotalGoalsConceded int    `json:"total_goals_conceded"`
	CleanSheets        int    `json:"clean_sheets"`
	Form               string `json:"form"`
}
