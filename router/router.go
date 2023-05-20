package router

import (
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test")
}

func Configure() *gin.Engine {
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

	router.GET("/", getHello)

	return router
}
