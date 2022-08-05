package ginupload

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/upload/uploadstorage"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func Upload(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				panic(common.ErrInternal(err))
			}
		}(file)

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStorage(appCtx.GetMySQLConnection())
		service := uploadservice.
			c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
