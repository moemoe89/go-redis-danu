package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"practicing-redis-golang/controllers"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	r.GET("/ping",controllers.Ping)

	r.POST("/hset",controllers.HSet)
	r.POST("/set",controllers.Set)
	r.POST("/sadd",controllers.SAdd)

	r.GET("/get/:key",controllers.Get)
	r.GET("/sadd/:key",controllers.GetSAdd)
	r.GET("/hset/:key/:field",controllers.GetHSet)
	r.DELETE("/delete/:key",controllers.Delete)

	r.Run()
}