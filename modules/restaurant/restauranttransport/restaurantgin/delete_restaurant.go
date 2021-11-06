package restaurantgin

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantbiz"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {

	 	id ,err := strconv.Atoi( c.Param("id"))
	//	uid ,err := common.FromBase58(c.Param("id"))
		if err!= nil{
			c.JSON(400 , gin.H{"error" : err.Error()})
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil{
			c.JSON(500 , gin.H{"error": err.Error()})
			return
		}

		c.JSON(200 , common.SimpleSucessResponse(true))
	}
}