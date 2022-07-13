package ginrestaurant

import (
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantservice"
	"food_delivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStorage(db)
		service := restaurantservice.NewCreateRestaurantService(store)

		if err := service.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}