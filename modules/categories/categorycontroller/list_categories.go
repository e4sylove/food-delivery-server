package categorycontroller

import (
	"food_delivery/common"
	"food_delivery/components/appctx"
	"food_delivery/modules/categories/categoryservice"
	"food_delivery/modules/categories/categorystorage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListCategories(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		categoriesStore := categorystorage.NewSQLStorage(appCtx.GetMySQLConnection())
		categoriesService := categoryservice.NewCategoriesService(categoriesStore)

		result, err := categoriesService.ListCategories(c.Request.Context())

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}