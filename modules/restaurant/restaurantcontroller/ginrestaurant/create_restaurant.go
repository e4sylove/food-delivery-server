package ginrestaurant

import (
	common2 "food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		db := appCtx.GetMySQLConnection()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common2.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStorage(db)
		service := restaurantservice.NewCreateRestaurantService(store)

		if err := service.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common2.SimpleSuccessResponse(data))
	}
}
