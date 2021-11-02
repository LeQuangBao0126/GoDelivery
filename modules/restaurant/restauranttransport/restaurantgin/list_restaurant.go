package restaurantgin

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantbiz"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		var paging common.Paging

		if err := c.ShouldBind(&filter) ; err!= nil{
			c.JSON(400 , gin.H{ "error" : "Parse data failure"})
			return
		}

		if err := c.ShouldBind(&paging) ; err!= nil{
			c.JSON(400 , gin.H{ "error" : "Parse data failure"})
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz( store)

		result ,err := biz.ListRestaurant(c.Request.Context(),&filter ,&paging)
		if err  != nil{
			c.JSON(400 , gin.H{ "error" :  err.Error()})
			return
		}

		c.JSON(201 , common.NewSucessResponse(result,paging,filter))
	}
}