package ginrestaurant

import (
	"food_delivery/components/appctx"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant (ctx appctx.AppContext) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := restaurantstorage.NewSQLStorage(ctx.GetMySQLConnection())
		service := restaurantservice.NewDeleteRestaurantService(storage)

		if err := service.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status": 1,
		})
	}
}