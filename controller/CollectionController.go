package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"sharephoto/model"
	"sharephoto/utils/mongodb"
	"time"
)

func CollectionController(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	var event = new(model.Event)

	fmt.Println("data base name", env.Get("DATABASE_NAME"))

	bs, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Collection", bson.M{
		"uuid": uuid,
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(bs) == 0 {
		//create collection

		bs2, err := bson.Marshal(model.Event{
			UUID:          uuid,
			StartDateTime: time.Now().Format(time.RFC3339),
			EndDateTime:   time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		})
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = mongodb.DB.InsertOne(env.Get("DATABASE_NAME"), "Collection", bs2)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = bson.Unmarshal(bs2, event)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
				"data":  "error unmarshal",
			})
			return
		}

	} else {
		// check if collection is expired
		// if expired, delete collection

		data, err := bson.Marshal(bs[0])
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = bson.Unmarshal(data, event)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		EndDate, err := time.Parse(time.RFC3339, event.EndDateTime)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if time.Now().After(EndDate) {
			_, err := mongodb.DB.DeleteOne(env.Get("DATABASE_NAME"), "Collection", bson.M{
				"uuid": uuid,
			})
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.Redirect(302, "/"+uuid)
			return
		}
	}

	ctx.HTML(200, "index.html", gin.H{
		"uuid":  uuid,
		"event": event,
	})
}
