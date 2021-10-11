package controller

import (
	"minesweeper-api/internal/service"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	Service service.SearchServiceInterface
}

func NewSearchController(s service.SearchServiceInterface) SearchController {
	return SearchController{Service: s}
}

func (controller SearchController) FindByGameId(c *gin.Context) {
	id := c.Param("id")
	response, err := controller.Service.FindByGameId(id)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"type": err.Type, "message": err.Message})
		return
	}

	c.JSON(200, response)
}
