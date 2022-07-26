package ginuser

import (
	"food_delivery/components/appctx"
	"food_delivery/components/hasher"
	"food_delivery/modules/common"
	"food_delivery/modules/user/usermodel"
	"food_delivery/modules/user/userservice"
	"food_delivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMySQLConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStorage(db)
		md5 := hasher.NewMd5Hash()
		service := userservice.NewRegisterService(store, md5)

		if err := service.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}