package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practicing-redis-golang/models"
)

func SAdd(c *gin.Context){

	var sAddRequest models.SAddRequest
	c.BindJSON(&sAddRequest)

	key := sAddRequest.Key
	value := sAddRequest.Value

	err := models.SAdd(key.(string),value.(string))
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponse(c,http.StatusOK,"Data has been saved.",true)
}