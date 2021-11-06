package restaurantgin

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantbiz"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	restaurantlikestorage "ProjectDelivery/modules/restaurantlike/storage"
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
		likeStore := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())

		//store likestore co the lay tren redis hay mongo  để hỗ trợ tăng tải
		biz := restaurantbiz.NewListRestaurantBiz( store ,likeStore)

		result ,err := biz.ListRestaurant(c.Request.Context(),&filter ,&paging)
		if err  != nil{
			c.JSON(400 , gin.H{ "error" :  err.Error()})
			return
		}

		//for i := range result{
		//	result[i].Mask(false)
		//}

		c.JSON(201 , common.NewSucessResponse(result,paging,filter))
	}
}