package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		// id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStorage(appCtx.GetMySQLConnection())
		service := restaurantservice.NewGetRestaurantService(store)

		result, err := service.GetRestaurantService(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		
		result.Mask(false)
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
