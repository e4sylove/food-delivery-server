package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStorage(appCtx.GetMySQLConnection())
		service := restaurantservice.NewGetRestaurantService(store)

		result, err := service.GetRestaurantService(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
