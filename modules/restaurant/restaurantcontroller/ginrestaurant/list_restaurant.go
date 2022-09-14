package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantrepo"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"food_delivery/modules/restaurantlike/restaurantlikestorage"
	"food_delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		restaurantStore := restaurantstorage.NewSQLStorage(appCtx.GetMySQLConnection())
		likeStore := restaurantlikestorage.NewSQLStorage(appCtx.GetMySQLConnection())
		userStore := userstorage.NewSQLStorage(appCtx.GetMySQLConnection())
		// userStore :=  remoteapi.NewUserAPI("http://localhost:3000")

		repository := restaurantrepo.NewListRestaurantRepo(restaurantStore, userStore, likeStore)
		service := restaurantservice.NewListRestaurantService(repository)

		result, err := service.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
