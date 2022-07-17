package ginrestaurant

import (
	"food_delivery/components"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateRestaurant(appCtx components.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}
		
		store := restaurantstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		service := restaurantservice.NewCreateRestaurantService(store)

		if err := service.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}