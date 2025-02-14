package models

import "time"

// ClubSeasonDataJSONRequest represents the request for the endpoint /club/season
type ClubSeasonDataJSONRequest struct {
	League string         `json:"league"`
	Data   []ModelFixture `json:"data"`
}

// ModelFixture represents a game between two teams
type ModelFixture struct {
	TeamOne string `json:"teamOne"`
	TeamTwo string `json:"teamTwo"`
}

// ClubSeasonDataJSONResponse response for endpoint /club/season
type ClubSeasonDataJSONResponse struct {
	HistoricalData []HistoricalData `json:"data"`
}

// HistoricalData respresnt the hitorical games between two games
type HistoricalData struct {
	TeamsID  string        `json:"id"`
	TeamOne  string        `json:"teamOne"`
	TeamTwo  string        `json:"teamTwo"`
	Fixtures []ClubFixture `json:"fixtures"`
}

// ClubFixture represents the details of a game between two teams
type ClubFixture struct {
	Div              string    `json:"div"`
	Date             time.Time `json:"date"`
	HomeTeam         string    `json:"homeTeam"`
	AwayTeam         string    `json:"awayTeam"`
	FullTimeHomeGoal string    `json:"fullTimeHomeGoal"`
	FullTimeAwayGoal string    `json:"fullTimeAwayGoal"`
	Winner           string    `json:"winner"`
	HalfTimeHomeGoal string    `json:"halfTimeHomeGoal"`
	HalfTimeAwayGoal string    `json:"halfTimeAwayGoal"`
	HalfTimeWinner   string    `json:"halfTimeWinner"`
	Attendance       string    `json:"attendance"`
	Referee          string    `json:"referee"`
	Season           int32     `json:"season"`
}

// WorldCupDataJSON represents all world cup data teams
type WorldCupDataJSON struct {
	TotalGoals          int32  `json:"total_goals"`
	HalfTimeGoals       int32  `json:"half_time_goals"`
	TeamGames           []Team `json:"team_games"`
	MostSuccessfulTeams []Team `json:"most_successful_teams"`
}

// Team represents a world cup participated team
type Team struct {
	Nation   string `json:"nation"`
	Amount   int    `json:"amount"`
	Position int    `json:"position"`
}

type TransferOfTeamRequest struct {
	Team    string   `json:"team"`
	Seasons []string `json:"seasons"`
	Window  string   `json:"window"`
}

type TransferOfTeamResponse struct {
	TransfersIn  []Transfer `json:"in"`
	TransfersOut []Transfer `json:"out"`
	BackFromLoad []Transfer `json:"back_from_loan"`
}

type Transfer struct {
	UUID             string  `json:"uuid"`
	ClubName         string  `json:"club_name"`
	PlayerName       string  `json:"player_name"`
	Age              int32   `json:"age"`
	Position         string  `json:"position"`
	PositionUUID     string  `json:"position_UUID"`
	ClubInvolvedName string  `json:"club_involved_name"`
	Fee              string  `json:"fee"`
	TransferMovement string  `json:"transfer_movement"`
	TransferPeriod   string  `json:"transfer_period"`
	FeeCleaned       float64 `json:"fee_cleaned"`
	LeagueName       string  `json:"league_name"`
	LeagueUUID       string  `json:"league_UUID"`
	Year             int32   `json:"year"`
	Season           string  `json:"season"`
}

type TransferMetadataResponse struct {
	Position []string         `json:"position"`
	Season   []string         `json:"season"`
	League   []LeagueMetadata `json:"leagues"`
}

type LeagueMetadata struct {
	Name  string   `json:"name"`
	Teams []string `json:"teams"`
}
