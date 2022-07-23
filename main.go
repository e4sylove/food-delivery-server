package main

import (
	"food_delivery/components"
	"food_delivery/helpers"
	"food_delivery/middleware"
	"food_delivery/modules/restaurant/restaurantcontroller/ginrestaurant"
	"food_delivery/modules/user/usercontroller/ginuser"
	"log"

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
	appCtx := components.NewAppContext(db)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.POST("/register", ginuser.Register(appCtx))

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run(`:8080`);
}