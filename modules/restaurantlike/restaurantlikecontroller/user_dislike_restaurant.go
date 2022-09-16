package restaurantlikecontroller

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurantlike/restaurantlikeservice"
	"food_delivery/modules/restaurantlike/restaurantlikestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UseDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMySQLConnection()

		store := restaurantlikestorage.NewSQLStorage(db)
		service := restaurantlikeservice.NewUserDislikeRestaurantService(store)

		err = service.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}