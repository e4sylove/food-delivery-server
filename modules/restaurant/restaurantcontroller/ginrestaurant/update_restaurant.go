package ginrestaurant

import (
	"food_delivery/components"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantservice"
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

		var data restaurantmodel.RestaurantUpdate
		
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		service := restaurantservice.NewUpdateRestaurantService(store)

		if err := service.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}