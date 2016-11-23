package controllers

import (
	"github.com/gin-gonic/gin"
	"practicing-redis-golang/models"
	"net/http"
)

func Get(c *gin.Context){

	key := c.Param("key")

	result,err := models.Get(key)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"Success get data.",true,result)
}

func GetSAdd(c *gin.Context){

	key := c.Param("key")

	result,err := models.GetSadd(key)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"Success get data.",true,result)
}

func GetHSet(c *gin.Context){

	key := c.Param("key")
	field := c.Param("field")

	result,err := models.GetHSet(key,field)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"Success get data.",true,result)
}