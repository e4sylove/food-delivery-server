package ginrestaurant

import (
	"food_delivery/components"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant (ctx components.AppContext) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		storage := restaurantstorage.NewSQLStorage(ctx.GetMySQLConnection())
		service := restaurantservice.NewDeleteRestaurantService(storage)

		if err := service.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status": 1,
		})
	}
}