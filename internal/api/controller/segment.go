package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaochengfu/go-id/internal/api/service"
	"net/http"
)

func Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		businessKey := c.Param("key")
		idModel := service.NewIdSequence()
		id, _ := idModel.GetId(businessKey)
		c.JSON(http.StatusOK, id)
	}
}
