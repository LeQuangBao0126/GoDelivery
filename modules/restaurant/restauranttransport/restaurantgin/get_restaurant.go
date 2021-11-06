package restaurantgin

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantbiz"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurantById(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {

		//uid ,err := common.FromBase58(c.Param("id"))
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			 //c.JSON(400, gin.H{"error":err.Error()})
			 //return
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(),id)

		if err != nil {
			c.JSON(400, err)
			return
		}
		//data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data))
	}
}

//phut 49 section5