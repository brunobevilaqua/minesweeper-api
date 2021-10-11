package service

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
	"minesweeper-api/internal/repository"
	"minesweeper-api/pkg/errors"
	"time"

	"github.com/google/uuid"
)

const (
	CLICK_ACTION string = "click"
	FLAG_ACTION  string = "flag"
)

type MaintenanceServiceInterface interface {
	PerformAction(id string, d dto.ActionRequest) (*dto.GameResponse, *errors.ApiError)
	CreateNewGame(d dto.CreateNewGameRequest) (*dto.GameResponse, *errors.ApiError)
}

type MaintenanceService struct {
	searchService SearchServiceInterface
	store         repository.Repository
}

func NewMaintenanceService(r repository.Repository, s SearchServiceInterface) MaintenanceService {
	return MaintenanceService{store: r, searchService: s}
}

func (s MaintenanceService) PerformAction(id string, request dto.ActionRequest) (*dto.GameResponse, *errors.ApiError) {
	game, board, err := s.searchService.FindGameAndBord(id)
	if err != nil {
		return nil, err
	}
	if game.Status == model.GAME_STATUS_LOST {
		return nil, errors.NewApiError(errors.LOST_GAME_ERROR)
	}

	// checks if game has a valid status...
	if game.Status != model.GAME_STATUS_PLAYING {
		return nil, errors.NewApiError(errors.GAME_ALREADY_ENDED)
	}

	switch request.Action {
	case CLICK_ACTION:
		return s.click(request.Row, request.Column, *game, *board)
	case FLAG_ACTION:
		return s.flag(request.Row, request.Column, *game, *board)
	default:
		return nil, errors.NewApiError(errors.INVALID_ACTION_ERROR)
	}
}

func (s MaintenanceService) flag(row, col int, g model.Game, b model.Board) (*dto.GameResponse, *errors.ApiError) {
	cell := &b.Grid[row][col]
	if cell.Status == model.CELL_CLICKED || cell.Status == model.CELL_EXPANDED {
		return nil, errors.NewApiError(errors.CELL_ALREADY_CLICKED_ERROR)
	} else if cell.Status == model.CELL_FLAGGED {
		b.Flags -= 1
		if cell.Mine { // removed flag from a mine
			b.MinesDiscovered -= 1
		}
		cell.Status = model.CELL_UNCLIKED
	} else {
		b.Flags += 1
		if cell.Mine { // added a flag to a correct mine
			b.MinesDiscovered += 1
		}
		cell.Status = model.CELL_FLAGGED
	}

	// checks if game ended, which means that all mines were discovered/flagged.
	if b.BoardEnded() {
		// game won, update game on redis
		g.EndTime = time.Now()
		g.Status = model.GAME_STATUS_WON
	}

	err := s.saveBoard(b)
	if err != nil {
		return nil, err
	}

	response := dto.NewGameResponse(
		dto.NewBoardDto(b.Rows, b.Cols, b.NumberOfMines, b.Clicks, b.MinesDiscovered, b.Grid, b.MinesPositions),
		dto.NewGameDto(g.Id, g.PlayerName, g.Status, g.StartTime, g.EndTime),
	)

	return &response, nil
}

func (s *MaintenanceService) click(row, col int, g model.Game, b model.Board) (*dto.GameResponse, *errors.ApiError) {
	boardUpdate := b // used in case of rollback.
	gameUpdate := g  // used in case of rollback.

	cell := &boardUpdate.Grid[row][col]
	// returns an error when cell already clicked
	if cell.Status == model.CELL_CLICKED || cell.Status == model.CELL_EXPANDED {
		return nil, errors.NewApiError(errors.CELL_ALREADY_CLICKED_ERROR)
	}

	boardUpdate.Clicks += 1

	// clicked on a mine, update game and board state on redis and return response showing mines
	if cell.Mine {
		gameUpdate.EndTime = time.Now()
		gameUpdate.Status = model.GAME_STATUS_LOST
		cell.Status = model.CELL_EXPLODED

		err := s.updateGameAndBoard(g, gameUpdate, b, boardUpdate)
		if err != nil {
			return nil, err
		}

		// returns response saying that user lost game and revelaing all the mines...
		response := dto.NewGameResponse(
			dto.NewBoardDto(boardUpdate.Rows, boardUpdate.Cols, boardUpdate.NumberOfMines, boardUpdate.Clicks, boardUpdate.MinesDiscovered, boardUpdate.Grid, boardUpdate.MinesPositions),
			dto.NewGameDto(gameUpdate.Id, gameUpdate.PlayerName, gameUpdate.Status, gameUpdate.StartTime, gameUpdate.EndTime),
		)
		return &response, nil
	}

	// when cell not clicked and not a mine, proceed expanding the other cells around.
	cell.Status = model.CELL_CLICKED
	cell.Evaluated = true

	if cell.NumberOfNearbyMines == 0 {
		// As the cell is empty (don't have mines nearby)...
		// Expanding the nearby Cells
		for _, cellPos := range cell.NearbyCells {
			cell := boardUpdate.Grid[cellPos.Col][cellPos.Row]
			boardUpdate.ExpandNearbyCell(cell)
		}
	}

	// updating board data on redis.
	s.saveBoard(boardUpdate)

	response := dto.NewGameResponse(
		dto.NewBoardDto(boardUpdate.Rows, boardUpdate.Cols, boardUpdate.NumberOfMines, boardUpdate.Clicks, boardUpdate.MinesDiscovered, boardUpdate.Grid, boardUpdate.MinesPositions),
		dto.NewGameDto(gameUpdate.Id, gameUpdate.PlayerName, gameUpdate.Status, gameUpdate.StartTime, gameUpdate.EndTime),
	)
	return &response, nil
}

func (s MaintenanceService) CreateNewGame(request dto.CreateNewGameRequest) (*dto.GameResponse, *errors.ApiError) {
	gameId := uuid.NewString()

	board := model.NewBoard(request.NumberOfRows,
		request.NumberOfColumns,
		request.NumberOfMines,
		gameId)
	game, err := model.NewGame(
		request.PlayerName,
		request.NumberOfRows,
		request.NumberOfColumns,
		request.NumberOfMines,
		gameId)

	if err != nil {
		return nil, err
	}

	if err := s.saveBoard(board); err != nil {
		return nil, err
	}

	if err := s.saveGame(*game); err != nil {
		s.store.DeleteBoardById(gameId)
		return nil, err
	} else {
		response := dto.NewGameResponse(
			dto.NewBoardDto(board.Rows, board.Cols, board.NumberOfMines, board.Clicks, board.MinesDiscovered, board.Grid, board.MinesPositions),
			dto.NewGameDto(game.Id, game.PlayerName, game.Status, game.StartTime, game.EndTime),
		)
		return &response, nil
	}
}

func (s MaintenanceService) saveGame(g model.Game) *errors.ApiError {
	err := s.store.SaveGame(g)
	if err != nil {
		return errors.NewApiError(errors.UPDATE_GAME_ERROR)
	}
	return nil
}

func (s MaintenanceService) saveBoard(b model.Board) *errors.ApiError {
	err := s.store.SaveBoard(b)
	if err != nil {
		return errors.NewApiError(errors.UPDATE_BOARD_ERROR)
	}
	return nil
}

func (s MaintenanceService) updateGameAndBoard(currentGame, gameUpdate model.Game, currentBoard, boardUpdate model.Board) *errors.ApiError {
	err := s.saveGame(gameUpdate)
	if err != nil {
		return err
	}

	err = s.saveBoard(boardUpdate)
	if err != nil {
		// rollback to previous game state and return error.
		s.saveGame(currentGame)
		return err
	}

	return nil
}
