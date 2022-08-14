package ginuser

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}