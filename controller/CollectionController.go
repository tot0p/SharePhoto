package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/SharePhoto/model"
	"github.com/tot0p/SharePhoto/utils/mongodb"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CollectionController(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	var event = new(model.Event)

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

	pictures, err := mongodb.DB.Find(env.Get("DATABASE_NAME"), "Picture", bson.M{
		"uuidevent": uuid,
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var picturesList []model.SimplePicture
	var top3 []model.SimplePicture

	for _, k := range pictures {
		var picture model.SimplePicture
		fmt.Println("k", k)
		if v, ok := k["uuid"]; ok {
			picture.UUID = v.(string)
		}
		if v, ok := k["like"]; ok {
			if v == nil {
				picture.Like = 0
			} else {
				picture.Like = len(v.(primitive.A))
			}
		}
		if v, ok := k["uuidevent"]; ok {
			picture.UUIDEvent = v.(string)
		}
		picturesList = append(picturesList, picture)
		if len(top3) < 3 {
			top3 = append(top3, picture)
		} else {
			for i, v := range top3 {
				if v.Like < picture.Like {
					temp := top3[i]
					top3[i] = picture
					picture = temp
				}
			}
		}
	}

	top := model.SimplePicture{}
	top1 := model.SimplePicture{}
	top2 := model.SimplePicture{}

	if len(top3) >= 3 {
		top = top3[0]
		top1 = top3[1]
		top2 = top3[2]
	}

	fmt.Println("top3", top3)

	ctx.HTML(200, "index.html", gin.H{
		"uuid":  uuid,
		"event": event,
		"list":  picturesList,
		"top1":  top,
		"top2":  top1,
		"top3":  top2,
	})
}
