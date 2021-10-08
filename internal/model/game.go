package model

import (
	"minesweeper-api/pkg/errors"
	"time"

	"github.com/google/uuid"
)

const (
	GAME_STATUS_PLAYING = "Playing..."
	GAME_STATUS_LOST    = "Game Over!"
	GAME_STATUS_WON     = "Won"
)

type Game struct {
	Player    *Player   `json:"Player"`
	Board     *Board    `json:"Board"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Clicks    int       `json:"clicks"`
	Status    string    `json:"status"`
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
		Id:        uuid.NewString(),
		Status:    GAME_STATUS_PLAYING,
	}, nil
}
