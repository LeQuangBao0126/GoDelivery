package restaurantgin

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantbiz"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {
		var data  restaurantmodel.RestaurantUpdate
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"error":err.Error()})
			return
		}
		if err2 := c.ShouldBind(&data) ; err2 != nil {
			c.JSON(400, gin.H{"error":err2.Error()})
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if	err := biz.UpdateRestaurant(c.Request.Context(),id,&data) ; err != nil {
			c.JSON(400, gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}