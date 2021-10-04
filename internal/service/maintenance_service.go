package service

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
	"minesweeper-api/internal/repository"
	"minesweeper-api/pkg/errors"
)

type MaintenanceServiceInterface interface {
	Click(d dto.ClickGameRequest) (*dto.Response, *errors.ApiError)
	CreateNewGame(d dto.CreateNewGameRequest) (*dto.Response, *errors.ApiError)
}

type MaintenanceService struct {
	store repository.Repository
}

func NewMaintenanceService(r repository.Repository) MaintenanceService {
	return MaintenanceService{store: r}
}

func (service MaintenanceService) Click(request dto.ClickGameRequest) (*dto.Response, *errors.ApiError) {

	return nil, nil
}

func (service MaintenanceService) CreateNewGame(request dto.CreateNewGameRequest) (*dto.Response, *errors.ApiError) {
	game, err := model.NewGame(request.PlayerName, request.NumberOfRows, request.NumberOfColumns, request.NumberOfMines)

	if err != nil {
		return nil, err
	}

	if newGame, err := service.store.Save(*game); err == nil {
		return dto.NewResponse(*newGame), nil
	} else {
		return nil, errors.NewApiError(errors.ErrorSavingEntity)
	}
}
