package service_test

import (
	"minesweeper-api/internal/dto"
	mocks "minesweeper-api/internal/mock"
	"minesweeper-api/internal/model"
	"minesweeper-api/internal/service"
	"minesweeper-api/pkg/errors"
	"testing"
)

func TestCreateGame(t *testing.T) {
	svc := service.MaintenanceService{
		Store: &mocks.RepositoryMock{
			OnSaveGame: func(g model.Game) error {
				return nil
			},
			OnSaveBoard: func(b model.Board) error {
				return nil
			},
		},
		SearchService: &mocks.SearchServiceMock{},
	}

	response, _ := svc.CreateNewGame(dto.CreateNewGameRequest{
		PlayerName:      "Bruno",
		NumberOfMines:   12,
		NumberOfColumns: 15,
		NumberOfRows:    15,
	})

	if response.Data.GameDto.PlayerName != "Bruno" &&
		response.Data.BoardDto.Columns != 15 &&
		response.Data.BoardDto.Rows != 15 {
		t.Fatalf("Game not created correclty")
	}
}

func TestPerformActionClick(t *testing.T) {
	svc := service.MaintenanceService{
		Store: &mocks.RepositoryMock{
			OnSaveGame: func(g model.Game) error {
				return nil
			},
			OnSaveBoard: func(b model.Board) error {
				return nil
			},
		},
		SearchService: &mocks.SearchServiceMock{
			OnFindGameAndBord: func(gameId string) (*model.Game, *model.Board, *errors.ApiError) {
				game := model.Game{
					PlayerName: "Bruno",
					Id:         "123",
				}

				board := model.Board{
					Rows:   3,
					Cols:   3,
					GameId: "123",
					Clicks: 0,
					Grid: []model.CellGrid{
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: true},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
					},
				}
				return &game, &board, nil
			},
		},
	}

	response, _ := svc.PerformAction("123", dto.ActionRequest{Action: "click", Row: 0, Column: 0})

	if response.Data.BoardDto.Clicks != 1 {
		t.Fatalf("Click not recoreded")
	}
}

func TestPerformActionClick_AlreadyClickedError(t *testing.T) {
	svc := service.MaintenanceService{
		Store: &mocks.RepositoryMock{
			OnSaveGame: func(g model.Game) error {
				return nil
			},
			OnSaveBoard: func(b model.Board) error {
				return nil
			},
		},
		SearchService: &mocks.SearchServiceMock{
			OnFindGameAndBord: func(gameId string) (*model.Game, *model.Board, *errors.ApiError) {
				game := model.Game{
					PlayerName: "Bruno",
					Id:         "123",
				}

				board := model.Board{
					Rows:   3,
					Cols:   3,
					GameId: "123",
					Clicks: 2,
					Grid: []model.CellGrid{
						{
							model.Cell{Evaluated: true, Status: model.CELL_CLICKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: true},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
					},
				}
				return &game, &board, nil
			},
		},
	}

	_, err := svc.PerformAction("123", dto.ActionRequest{Action: "click", Row: 0, Column: 0})

	if err == nil {
		t.Fatalf("should return already clicked error")
	}
}

func TestPerformActionFlag(t *testing.T) {
	svc := service.MaintenanceService{
		Store: &mocks.RepositoryMock{
			OnSaveGame: func(g model.Game) error {
				return nil
			},
			OnSaveBoard: func(b model.Board) error {
				return nil
			},
		},
		SearchService: &mocks.SearchServiceMock{
			OnFindGameAndBord: func(gameId string) (*model.Game, *model.Board, *errors.ApiError) {
				game := model.Game{
					PlayerName: "Bruno",
					Id:         "123",
				}

				board := model.Board{
					Rows:            3,
					Cols:            3,
					GameId:          "123",
					Clicks:          2,
					MinesDiscovered: 0,
					Grid: []model.CellGrid{
						{
							model.Cell{Evaluated: true, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: true},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: true},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
					},
				}
				return &game, &board, nil
			},
		},
	}

	response, err := svc.PerformAction("123", dto.ActionRequest{Action: "flag", Row: 0, Column: 0})

	if err != nil {
		t.Fatalf("should not return error, as the cell is not clicked")
	}

	if response.Data.BoardDto.DiscoveredMines != 1 {
		t.Fatalf("should increment the number of discovered mines")
	}
}

func TestPerformActionFlag_AlreadyClickedError(t *testing.T) {
	svc := service.MaintenanceService{
		Store: &mocks.RepositoryMock{
			OnSaveGame: func(g model.Game) error {
				return nil
			},
			OnSaveBoard: func(b model.Board) error {
				return nil
			},
		},
		SearchService: &mocks.SearchServiceMock{
			OnFindGameAndBord: func(gameId string) (*model.Game, *model.Board, *errors.ApiError) {
				game := model.Game{
					PlayerName: "Bruno",
					Id:         "123",
				}

				board := model.Board{
					Rows:   3,
					Cols:   3,
					GameId: "123",
					Clicks: 2,
					Grid: []model.CellGrid{
						{
							model.Cell{Evaluated: true, Status: model.CELL_CLICKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: true},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
						{
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 2, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 3, Mine: false},
							model.Cell{Evaluated: false, Status: model.CELL_UNCLIKED, NumberOfNearbyMines: 1, Mine: false},
						},
					},
				}
				return &game, &board, nil
			},
		},
	}

	_, err := svc.PerformAction("123", dto.ActionRequest{Action: "flag", Row: 0, Column: 0})

	if err == nil {
		t.Fatalf("should return already clicked error")
	}
}
