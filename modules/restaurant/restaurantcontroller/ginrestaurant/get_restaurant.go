package ginrestaurant

import (
	"food_delivery/components"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetRestaurant(appCtx components.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		serivce := restaurantservice.NewGetRestaurantService(store)
		
		result, err := serivce.GetRestaurantService(c.Request.Context(), id)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}