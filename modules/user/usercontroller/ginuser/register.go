package ginuser

import (
	"food_delivery/components"
	"food_delivery/modules/common"
	"food_delivery/modules/user/usermodel"
	"food_delivery/modules/user/userservice"
	"food_delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx components.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMySQLConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStorage(db)
		md5 := components.NewMd5Hash()
		biz := userservice.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}


		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}