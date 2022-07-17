package main

import (
	"food_delivery/components"
	"food_delivery/helpers"
	"food_delivery/modules/restaurant/restaurantcontroller/ginrestaurant"
	"food_delivery/modules/restaurant/restaurantmodel"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := helpers.GetDsn("MYSQL_CONNECTION")
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	if err := serve(db); err != nil {
		log.Fatalln(err)
	}
}


func serve(db *gorm.DB) error {
	
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := components.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})
			}

			if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}
				
			c.JSON(http.StatusOK, gin.H{"status": 1})
		})
	}

	return r.Run(`:8080`);
}