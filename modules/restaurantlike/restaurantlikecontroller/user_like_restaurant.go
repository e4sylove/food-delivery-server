package restaurantlikecontroller

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"
	"food_delivery/modules/restaurantlike/restaurantlikeservice"
	"food_delivery/modules/restaurantlike/restaurantlikestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		db := appCtx.GetMySQLConnection()

		store := restaurantlikestorage.NewSQLStorage(db)

		service := restaurantlikeservice.NewUserLikeRestaurantService(store)

		if err := service.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}