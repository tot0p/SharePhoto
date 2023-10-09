package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"sharephoto/controller"
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

	//load templates
	r.LoadHTMLGlob("src/templates/*")

	r.Static("/static", "./src/static")

	// Index
	r.GET("/", controller.IndexController)
	r.GET("/index", controller.IndexController)

	api := r.Group("/api")

	api.POST("/upload", controller.UploadApiController)

	//api.GET("/get/:uuid", controller.GetController)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(env.Get("PORT")); err != nil {
		panic(err)
	}
}
