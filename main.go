package main

import (
	"food_delivery/components/appctx"
	"food_delivery/helpers"
	"food_delivery/middleware"
	"food_delivery/modules/restaurant/restaurantcontroller/ginrestaurant"
	"food_delivery/modules/upload/uploadcontroller/ginupload"
	"food_delivery/modules/user/usercontroller/ginuser"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := helpers.GetDsn("MYSQL_CONNECTION")
	secretKey := helpers.GetSecretKey("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := serve(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func serve(db *gorm.DB, secretKey string) error {
	appCtx := appctx.NewAppContext(db, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.POST("/upload", ginupload.Upload(appCtx))

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run(`:8080`)
}
