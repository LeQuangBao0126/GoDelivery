package ginrestaurantlike

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	rstrtlikebiz "ProjectDelivery/modules/restaurantlike/biz"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	restaurantlikestorage "ProjectDelivery/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUsersLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter  restaurantlikemodel.Filter
		var paging common.Paging

		id,err:= strconv.Atoi(c.Param("id"))
	    filter.RestaurantId = id

		if err := c.ShouldBind(&paging) ; err!= nil{
			c.JSON(400 , gin.H{ "error" : "Parse data failure"})
			return
		}
		likeStore := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())

		biz := rstrtlikebiz.NewListUserLikeRestaurantBiz(likeStore)

		result ,err := biz.ListUsers(c.Request.Context(),&filter,&paging)
		if err  != nil{
			c.JSON(400 , gin.H{ "error" :  err.Error()})
			return
		}
		c.JSON(201 , common.NewSucessResponse(result,paging,filter))
	}
}





