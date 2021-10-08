package controller

import (
	"minesweeper-api/internal/service"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	service service.SearchServiceInterface
}

func NewSearchController(s service.SearchService) SearchController {
	return SearchController{service: s}
}

func (controller SearchController) FindByUserName(c *gin.Context) {
	name := c.Param("name")
	response, err := controller.service.FindByPlayerName(name)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"type": err.Type, "message": err.Message})
		return
	}

	c.JSON(200, response)
}

func (controller SearchController) FindByGameId(c *gin.Context) {
	id := c.Param("id")
	response, err := controller.service.FindByGameId(id)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"type": err.Type, "message": err.Message})
		return
	}

	c.JSON(200, response)
}
