package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/SharePhoto/utils/mongodb"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
)

func ImgController(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	uuidImg := ctx.Param("uuidImg")

	// get img info from db

	collection, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Collection", bson.M{
		"uuid": uuid,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(collection) == 0 {
		ctx.JSON(404, gin.H{
			"error": "collection not found",
		})
		return
	}

	// get img from db

	fmt.Println("uuidImg", uuidImg)

	img, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Picture", bson.M{
		"uuid": uuidImg,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if img exists
	if len(img) == 0 {
		ctx.JSON(404, gin.H{
			"error": "img not found",
		})
		return
	}

	if path, ok := img[0]["path"].(string); ok {
		ctx.File(path)
		return
	}

	ctx.JSON(500, gin.H{
		"error": "path not found",
	})
}
