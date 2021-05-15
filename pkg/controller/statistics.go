package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsResponse struct {
	ID       int    `json:"id" binding:"id"`
	Name     string `json:"name" db:"name"`
	Likes    int    `json:"likes" db:"count_likes"`
	Comments int    `json:"comments" db:"count_comments"`
}

func (c *ControllerMain) getStatistics(ctx *gin.Context) {
	result, err := c.services.GetStatistics()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, result)
}
