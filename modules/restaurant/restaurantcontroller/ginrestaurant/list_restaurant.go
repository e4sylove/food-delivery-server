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

func ListRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()
		
		storage := restaurantstorage.NewSQLStorage(appCtx.GetMainDBConnection())
		service := restaurantservice.NewListRestaurantService(storage)

		result, err := service.ListRestaurant(c.Request.Context(), &filter, &paging) 
		
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	
	}
}