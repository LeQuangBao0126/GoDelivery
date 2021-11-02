package  restaurantgin

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantbiz"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
			var data restaurantmodel.RestaurantCreate

			if err := c.ShouldBind(&data) ; err!= nil{
				 c.JSON(400 , gin.H{ "error" : "Parse data failure"})
				return
			}
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			biz := restaurantbiz.NewCreateRestaurantBiz( store)

			if err := biz.CreateRestaurant(c.Request.Context() , &data) ; err!= nil{
				c.JSON(400 , gin.H{ "error" :err.Error()})
				return
			}
			c.JSON(201 , common.SimpleSucessResponse(data))
	}
}
/*type storeFake struct {
}
func (storeFake) Create(ctx context.Context ,data * restaurantmodel.RestaurantCreate) error{
	return errors.New("Fake store ")
}*/