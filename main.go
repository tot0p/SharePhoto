package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/SharePhoto/controller"
	"github.com/tot0p/SharePhoto/utils"
	"github.com/tot0p/SharePhoto/utils/mongodb"
	"github.com/tot0p/env"
)

func init() {

	_ = env.Load()

	// create src/cdn folder if not exists
	utils.CreateDirIfNotExists("src/cdn")

	err := mongodb.NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
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
	r.GET("/:uuid/img/:uuidImg", controller.ImgController)
	r.GET("/:uuid/img/:uuidImg/like", controller.LikeController)

	//sessionManager
	r.POST("/fingerprint", controller.FingerPrintApiController)

	//api.GET("/get/:uuid", controller.GetController)

	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
