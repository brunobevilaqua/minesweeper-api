package service

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
	"minesweeper-api/internal/repository"
	"minesweeper-api/pkg/errors"
)

type SearchServiceInterface interface {
	FindByGameId(id string) (*dto.GameResponse, *errors.ApiError)
	FindBoardById(id string) (*model.Board, *errors.ApiError)
	FindGame(id string) (*model.Game, *errors.ApiError)
	FindGameAndBord(gameId string) (*model.Game, *model.Board, *errors.ApiError)
}

type SearchService struct {
	store repository.Repository
}

func NewSearchService(r repository.Repository) SearchService {
	return SearchService{store: r}
}

func (s SearchService) FindByGameId(id string) (*dto.GameResponse, *errors.ApiError) {
	if id == "" {
		return nil, errors.NewApiError(errors.INVALID_PARAMETER_ERROR)
	}

	game, err := s.store.FindGameById(id)

	if err != nil {
		return nil, errors.NewApiError(errors.NO_RECORDS_FOUND_ERROR)
	}

	board, err := s.store.FindBoardById(id)

	if err != nil {
		return nil, errors.NewApiError(errors.NO_RECORDS_FOUND_ERROR)
	}

	response := dto.NewGameResponse(
		dto.NewBoardDto(board.Rows, board.Cols, board.NumberOfMines, board.Clicks, board.MinesDiscovered, board.Grid, board.MinesPositions),
		dto.NewGameDto(game.Id, game.PlayerName, game.Status, game.StartTime, game.EndTime),
	)

	return &response, nil
}

func (s SearchService) FindBoardById(id string) (*model.Board, *errors.ApiError) {
	board, err := s.store.FindBoardById(id)
	if err != nil {
		apiError := errors.NewApiError(errors.CUSTOM_ERROR)
		apiError.Message = err.Error()
		return nil, apiError
	}
	if board == nil {
		apiError := errors.NewApiError(errors.CUSTOM_ERROR)
		apiError.Message = "Board not Found!"
		return nil, apiError
	}
	return board, nil
}

func (s SearchService) FindGame(id string) (*model.Game, *errors.ApiError) {
	game, err := s.store.FindGameById(id)
	if err != nil {
		apiError := errors.NewApiError(errors.CUSTOM_ERROR)
		apiError.Message = err.Error()
		return nil, apiError
	}
	if game == nil {
		apiError := errors.NewApiError(errors.CUSTOM_ERROR)
		apiError.Message = "Game not Found!"
		return nil, apiError
	}
	return game, nil
}

func (s SearchService) FindGameAndBord(gameId string) (*model.Game, *model.Board, *errors.ApiError) {
	game, err := s.FindGame(gameId)
	if err != nil || game == nil {
		return nil, nil, errors.NewApiError(errors.NO_RECORDS_FOUND_ERROR)
	}

	board, err := s.FindBoardById(gameId)
	if err != nil {
		return nil, nil, err
	}

	return game, board, nil
}
