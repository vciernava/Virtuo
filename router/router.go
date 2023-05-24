package router

import (
	"../environment"
	"./routes"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

func Configure() *gin.Engine {
	routes.NewInstall()

	gin.SetMode("release")

	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		log.WithFields(log.Fields{
			"client_ip":  params.ClientIP,
			"status":     params.StatusCode,
			"latency":    params.Latency,
			"request_id": params.Keys["request_id"],
		}).Debugf("%s %s", params.MethodColor()+params.Method+params.ResetColor(), params.Path)

		return ""
	}))

	router.POST("/image/pull", environment.PullImage)
	router.GET("/servers", routes.GetServers)
	router.POST("/server/create", routes.CreateServer)
	router.PATCH("/server/start", routes.StartServer)
	router.PATCH("/server/stop", routes.StopServer)
	router.DELETE("/server/delete", routes.DeleteServer)

	return router
}
