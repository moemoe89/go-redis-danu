package controllers

import (
	"github.com/gin-gonic/gin"
	"practicing-redis-golang/models"
	"net/http"
)

func Delete(c *gin.Context){

	key := c.Param("key")

	err := models.Delete(key)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponse(c,http.StatusOK,"Success delete.",true)
}