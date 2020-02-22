//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/practicing-redis-golang/api"
	mw "github.com/moemoe89/practicing-redis-golang/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(log *logrus.Entry, svc ap.Service) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mw.CORS)
	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	ctrl := ap.NewCtrl(log, svc)

	r.POST("/set", ctrl.Set)
	r.POST("/sadd", ctrl.SAdd)
	r.POST("/hset", ctrl.HSet)

	r.GET("/get/:key", ctrl.Get)
	r.GET("/sadd/:key", ctrl.GetSAdd)
	r.GET("/hset/:key/:field", ctrl.GetHSet)
	r.DELETE("/delete/:key", ctrl.Delete)

	return r
}
