package router

import (
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

	router.GET("/", getHello)

	return router
}
