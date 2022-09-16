package main

import (
	"food_delivery/components/appctx"
	"food_delivery/components/cache"
	"food_delivery/helpers"
	"log"

	"food_delivery/components/uploadprovider"
	"food_delivery/middleware"
	"food_delivery/modules/categories/categorycontroller"
	"food_delivery/modules/restaurant/restaurantcontroller/ginrestaurant"
	"food_delivery/modules/restaurantlike/restaurantlikecontroller"
	"food_delivery/modules/user/usercontroller/ginuser"
	"food_delivery/modules/user/usercontroller/internalapi"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

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

	redis := cache.NewRedisClient();

	db = db.Debug()
	
	if err := serve(db, secretKey, s3Provider, redis); err != nil {
		log.Fatalln(err)
	}

	// service := goservice.New(
	// 	goservice.WithName("food-delivery"),
	// 	goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
	// 	goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.JWTProvider)),
	// )
}

func serve(db *gorm.DB, 
	secretKey string, 
	uploadProvider uploadprovider.UploadProvider, 
	redis *redis.Client) error {
		
	appCtx := appctx.NewAppContext(db, secretKey, uploadProvider, redis)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	// v1.POST("/upload", ginupload.Upload(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	
	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		restaurants.POST("/:id/like", restaurantlikecontroller.UserLikeRestaurant(appCtx))

		// restaurants.POST("/:id/like", handlers ...gin.HandlerFunc)
	}

	internal := r.Group("/internal")
	{
		internal.POST("/get-users-by-ids", internalapi.GetUserById(appCtx))
	}

	categories := v1.Group("/categories")
	{
		categories.GET("/", categorycontroller.ListCategories(appCtx))
	}

	return r.Run(`:3000`)
}
