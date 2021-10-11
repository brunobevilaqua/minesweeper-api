package model

import "minesweeper-api/pkg/errors"

type Player struct {
	Name string `json:"name"`
}

// func newPlayer(n string) (Player, errors.ApiError) {
func NewPlayer(n string) (*Player, *errors.ApiError) {
	if n != "" {
		return &Player{Name: n}, nil
	} else {
		return nil, errors.NewApiError(errors.INVALID_USER_NAME_ERROR)
	}
}
