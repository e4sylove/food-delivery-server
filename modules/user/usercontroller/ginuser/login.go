package ginuser

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/components/hasher"
	"food_delivery/components/tokenprovider/jwt"
	"food_delivery/modules/user/usermodel"
	"food_delivery/modules/user/userservice"
	"food_delivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLogin usermodel.UserLogin

		if err := c.ShouldBind(&userLogin); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		md5 := hasher.NewMd5Hash()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecret())
		db := appCtx.GetMySQLConnection()
		store := userstorage.NewSQLStorage(db)
		service := userservice.NewLoginService(store, tokenProvider, md5, 60*60*24*30)

		account, err := service.Login(c.Request.Context(), &userLogin)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
