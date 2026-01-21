package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// Define a simple GET endpoint under /v1
		v1.GET("/ping", pong)
		v1.POST("/ping", pong)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", pong)
		v2.POST("/ping", pong)
		v2.PUT("/ping", pong)
	}

	return r
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}