package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
)

func init() {

	err := env.Load()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
