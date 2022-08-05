package main

import (
	"food_delivery/components/appctx"
	"food_delivery/components/uploadprovider"
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

	s3BucketName := helpers.GetS3BucketName("S3BUCKET_NAME")
	s3Region := helpers.GetS3Region("S3REGION")
	s3APIKey := helpers.GetS3APIKey("S3API_KEY")
	s3SecretKey := helpers.GetSecretKey("S3SECRET_KEY")
	s3Domain := helpers.GetS3Domain("S3DOMAIN")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := serve(db, secretKey, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func serve(db *gorm.DB, secretKey string, uploadProvider uploadprovider.UploadProvider) error {
	appCtx := appctx.NewAppContext(db, secretKey, uploadProvider)

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
