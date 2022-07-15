package ginrestaurant

import (
	"food_delivery/components"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func UpdateRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		var data restaurantmodel.Restaurant
		
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		// service := restaurantservice.NewUpdateRestaurantService(store)

	}
}