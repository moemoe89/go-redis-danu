package controllers

import (
	"github.com/gin-gonic/gin"
	"practicing-redis-golang/models"
)

func HSet(c *gin.Context){

	var hSetRequest models.HSetRequest
	c.BindJSON(&hSetRequest)

	key := hSetRequest.Key
	field := hSetRequest.Field
	value := hSetRequest.Value

	httpStatus,message,err := models.HSet(key.(string),field.(string),value.(string))
	if err != nil {
		JSONResponse(c,httpStatus,err.Error(),false)
		return
	}

	JSONResponse(c,httpStatus,message,true)
}