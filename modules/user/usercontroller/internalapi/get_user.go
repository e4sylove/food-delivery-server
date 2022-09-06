package internalapi

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetUserById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			Ids []int `json:"ids"`
		}

		if err := c.ShouldBind(&param); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMySQLConnection()
		store := userstorage.NewSQLStorage(db)

		result, err := store.GetUsers(c.Request.Context(), param.Ids)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(common.DbTypeUser)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
