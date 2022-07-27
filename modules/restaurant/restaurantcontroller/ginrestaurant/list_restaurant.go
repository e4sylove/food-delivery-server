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

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common2.ErrInvalidRequest(err))
		}

		var paging common2.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common2.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		storage := restaurantstorage.NewSQLStorage(appCtx.GetMySQLConnection())
		service := restaurantservice.NewListRestaurantService(storage)

		result, err := service.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common2.NewSuccessResponse(result, paging, filter))

	}
}
