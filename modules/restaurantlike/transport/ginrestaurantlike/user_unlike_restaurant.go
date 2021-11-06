package ginrestaurantlike

import (
	"ProjectDelivery/component"
	rstrtlikebiz "ProjectDelivery/modules/restaurantlike/biz"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	restaurantlikestorage "ProjectDelivery/modules/restaurantlike/storage"
	"ProjectDelivery/modules/user/usermodel"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserUnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {


		id,_:= strconv.Atoi(c.Param("id"))
		requester := c.MustGet("user").(*usermodel.User)
		data := restaurantlikemodel.Like{
			RestaurantId:  id,
			UserId: requester.Id,
		}

		unlikeStore := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		//descStore := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())

		biz := rstrtlikebiz.NewUserUnLikeRestaurantBiz(unlikeStore,appCtx.GetPubsub())
		if err := biz.UnlikeRestaurant(c.Request.Context(),&data) ; err!= nil{
			c.JSON(400 , err)
			return
		}

		c.JSON(201 , gin.H{"ok":true})
	}
}
