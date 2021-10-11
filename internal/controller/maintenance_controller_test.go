package controller_test

import (
	"minesweeper-api/internal/controller"
	"minesweeper-api/internal/dto"
	mocks "minesweeper-api/internal/mock"
	"minesweeper-api/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateNewGame(t *testing.T) {
	controller := controller.MaintenanceController{
		Service: &mocks.MaintenanceServiceMock{
			OnCreateNewGame: func(d dto.CreateNewGameRequest) (*dto.GameResponse, *errors.ApiError) {
				gDto := dto.GameDto{
					Id:         "123",
					PlayerName: "Bruno",
				}
				bDto := dto.BoardDto{
					Columns:       10,
					Rows:          10,
					Clicks:        0,
					NumberOfMines: 10,
				}
				response := dto.GameResponse{}
				response.Data.BoardDto = bDto
				response.Data.GameDto = gDto
				return &response, nil
			},
		},
	}

	r := gin.Default()
	r.POST("/games", func(c *gin.Context) {
		controller.CreateNewGame(*c)
	})
	r.Run(":8080")

	data := `{"playerName": "Bruno Bevilaqua","numberOfMines": 8,"numberOfColumns": 5,"numberOfRows": 5}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/games", strings.NewReader(data))
	r.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Fatalf("Shoud return 201 when creating a new game")
	}
}

func TestCreateNewGame_InvalidRequest(t *testing.T) {
	controller := controller.MaintenanceController{
		Service: &mocks.MaintenanceServiceMock{
			OnCreateNewGame: func(d dto.CreateNewGameRequest) (*dto.GameResponse, *errors.ApiError) {
				gDto := dto.GameDto{
					Id:         "123",
					PlayerName: "Bruno",
				}
				bDto := dto.BoardDto{
					Columns:       10,
					Rows:          10,
					Clicks:        0,
					NumberOfMines: 10,
				}
				response := dto.GameResponse{}
				response.Data.BoardDto = bDto
				response.Data.GameDto = gDto
				return &response, nil
			},
		},
	}

	r := gin.Default()
	r.POST("/games", func(c *gin.Context) {
		controller.CreateNewGame(*c)
	})
	r.Run(":8080")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/games", nil)
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fatalf("Shoud return 201 when creating a new game")
	}
}

func TestClick(t *testing.T) {
	controller := controller.MaintenanceController{
		Service: &mocks.MaintenanceServiceMock{
			OnAction: func(id string, d dto.ActionRequest) (*dto.GameResponse, *errors.ApiError) {
				return &dto.GameResponse{}, nil
			},
		},
	}

	r := gin.Default()
	r.PUT("/games", func(c *gin.Context) {
		controller.Click(*c)
	})
	r.Run(":8080")

	data := `{
		"action": "click", 
		"row": 6,
		"column": 8 
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/games", strings.NewReader(data))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Shoud return 200 when creating a new game")
	}
}

func TestClick_InvalidRequest(t *testing.T) {
	controller := controller.MaintenanceController{
		Service: &mocks.MaintenanceServiceMock{
			OnAction: func(id string, d dto.ActionRequest) (*dto.GameResponse, *errors.ApiError) {
				return &dto.GameResponse{}, nil
			},
		},
	}

	r := gin.Default()
	r.PUT("/games", func(c *gin.Context) {
		controller.Click(*c)
	})
	r.Run(":8080")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/games", nil)
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fatalf("Shoud return 200 when creating a new game")
	}
}
