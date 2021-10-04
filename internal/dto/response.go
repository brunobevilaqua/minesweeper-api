package dto

import "minesweeper-api/internal/model"

type Response struct {
	GameId         string `json:"gameId"`
	PlayerName     string `json:"playerName"`
	NumberOfClicks int    `json:"numberOfClicks"`
	Status         string `json:"status"`
	Points         int    `json:"points"`
}

func NewResponse(g model.Game) *Response {
	r := Response{
		PlayerName:     g.Player.Name,
		GameId:         g.Id,
		NumberOfClicks: g.Clicks,
		Status:         g.GetStatus(),
		Points:         g.Points,
	}

	return &r
}
