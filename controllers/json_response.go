package controllers

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context,httpStatus int,message string,status bool){
	c.IndentedJSON(httpStatus, gin.H{
		"message": message,
		"status": status,
	})
}

func JSONResponseData(c *gin.Context,httpStatus int,message string,status bool,data interface{}) {
	c.IndentedJSON(httpStatus, gin.H{
		"data": data,
		"message": message,
		"status": status,
	})
}