package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practicing-redis-golang/models"
)

func Set(c *gin.Context){

	var setRequest models.SetRequest
	c.BindJSON(&setRequest)

	key := setRequest.Key
	value := setRequest.Value

	err := models.Set(key.(string),value.(string))
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponse(c,http.StatusOK,"Data has been saved.",true)
}