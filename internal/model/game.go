package model

import (
	"minesweeper-api/pkg/errors"
	"time"
)

type Game struct {
	Player    *Player   `json:"Player"`
	Board     *Board    `json:"Board"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Clicks    int       `json:"clicks"`
	Lost      bool      `json:"lost"`
	Won       bool      `json:"won"`
	Points    int       `json:"points"`
	Id        string    `json:"id"`
}

func NewGame(playerName string, rows, cols, mines int) (*Game, *errors.ApiError) {
	player, err := NewPlayer(playerName)

	if err != nil {
		return nil, err
	}

	board := NewBoard(rows, cols, mines)

	return &Game{
		Player:    player,
		Board:     board,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Clicks:    0,
		// TODO generate GAME ID
	}, nil
}

func (g Game) GetStatus() string {
	if !g.Won && !g.Lost {
		return "Still Playing..."
	}
	if g.Won {
		return "Won!"
	} else {
		return "Game Over – You Lose!"
	}
}
