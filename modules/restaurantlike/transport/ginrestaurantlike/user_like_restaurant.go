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

func PostUsersLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {


		id,_:= strconv.Atoi(c.Param("id"))
		requester := c.MustGet("user").(*usermodel.User)
	 	data := restaurantlikemodel.Like{
			 RestaurantId:  id,
			 UserId: requester.Id,
		}

		likeStore := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		//incStore := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())

		biz := rstrtlikebiz.NewUserLikeRestaurantBiz(likeStore,appCtx.GetPubsub())

		if  err := biz.LikeRestaurant(c.Request.Context(),&data) ;err  != nil{
			c.JSON(400 , gin.H{ "error" :  err.Error()})
			return
		}
		c.JSON(201 , gin.H{"ok":true})
	}
}



