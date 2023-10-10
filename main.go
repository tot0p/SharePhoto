package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"sharephoto/controller"
	"sharephoto/utils"
)

func init() {

	err := env.Load()
	if err != nil {
		panic(err)
	}

	// create src/cdn folder if not exists
	utils.CreateDirIfNotExists("src/cdn")

	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.Default()

	//load templates
	r.LoadHTMLGlob("src/templates/*")

	r.Static("/static", "./src/static")

	// Index
	r.GET("/", controller.IndexController)

	//collection
	r.GET("/:uuid", controller.CollectionController)
	r.POST("/:uuid/upload", controller.UploadPostController)

	//upload
	//r.GET("/upload", controller.UploadController)
	//r.POST("/upload", controller.UploadPostController)

	//sessionManager
	r.POST("/fingerprint", controller.FingerPrintApiController)

	//api.GET("/get/:uuid", controller.GetController)

	if err := r.Run(env.Get("PORT")); err != nil {
		panic(err)
	}
}
