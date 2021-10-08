package controller

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MaintenanceController struct {
	Service service.MaintenanceServiceInterface
}

func NewMaintenanceController(s service.MaintenanceServiceInterface) MaintenanceController {
	return MaintenanceController{Service: s}
}

func (controller MaintenanceController) CreateNewGame(c gin.Context) {
	request := dto.CreateNewGameRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := controller.Service.CreateNewGame(request)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"type": err.Type, "message": err.Message})
		return
	}

	c.JSON(200, response)
}

func (controller MaintenanceController) Click(c gin.Context) {
	request := dto.ClickGameRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	response, err := controller.Service.Click(request)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"type": err.Type, "message": err.Message})
		return
	}

	c.JSON(200, response)
}
