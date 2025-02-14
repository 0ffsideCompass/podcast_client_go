package models

import "time"

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	League    string    `json:"league"`
	HTML      string    `json:"html"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []string  `json:"tags"`
	IMG       string    `json:"img"`
	Data      Data      `json:"data"`
}

type Articles struct {
	Data []Article `json:"data"`
}

type Data struct {
	TeamOne      string    `json:"teamOne"`
	TeamTwo      string    `json:"teamTwo"`
	TeamOneStats TeamStats `json:"teamOneStats"`
	TeamTwoStats TeamStats `json:"teamTwoStats"`
	LastGame     LastGame  `json:"lastGame"`
}

// TeamStats represents the stats of a team
type TeamStats struct {
	WinsThisSeason     int    `json:"winsThisSeason"`
	LosesThisSeason    int    `json:"lossesThisSeason"`
	DrawsThisSeason    int    `json:"drawsThisSeason"`
	TotalGames         int    `json:"totalGames"`
	TotalGoals         int    `json:"totalGoals"`
	TotalGoalsConceded int    `json:"totalGoalsConceded"`
	CleanSheets        int    `json:"cleanSheets"`
	Form               string `json:"form"`
}

type ArticleImportanceList struct {
	PositionOne   ArticleImportance `json:"positionOne"`
	PositionTwo   ArticleImportance `json:"positionTwo"`
	PositionThree ArticleImportance `json:"positionThree"`
	PositionFour  ArticleImportance `json:"positionFour"`
	PositionFive  ArticleImportance `json:"positionFive"`
	PositionSix   ArticleImportance `json:"positionSix"`
	PositionSeven ArticleImportance `json:"positionSeven"`
}

type ArticleImportance struct {
	ID    string `bson:"article_id"`
	Title string `bson:"article_title"`
}
