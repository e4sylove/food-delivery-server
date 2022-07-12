package main

import (
	"food_delivery/helpers"
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

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", func(c *gin.Context) {
			
			var data restaurantmodel.Restaurant

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}
			
			if err := db.Create(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})
			
				return
			}
			
			c.JSON(http.StatusOK, data)
		})

		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			var data restaurantmodel.Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})

				return 
			}

			c.JSON(http.StatusOK, data)
		})

		restaurants.GET("", func(c *gin.Context) {
			var data []restaurantmodel.Restaurant

			type Filter struct {
				CityID int `json:"city_id" form:"city_id"`
			}
			
			var filter Filter
			
			if err := c.ShouldBind(&filter); err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			newDb := db
			if filter.CityID > 0 {
				newDb = db.Where("city_id = ?", filter.CityID)
			}

			if err := newDb.Find(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}
			
			c.JSON(http.StatusOK, data)
		})
	}

	return r.Run(`:8080`);
}