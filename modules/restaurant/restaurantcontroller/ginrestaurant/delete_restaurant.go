package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(ctx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := restaurantstorage.NewSQLStorage(ctx.GetMySQLConnection())
		service := restaurantservice.NewDeleteRestaurantService(storage)

		if err := service.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status": 1,
		})
	}
}
