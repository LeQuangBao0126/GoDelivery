package ginuser

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/component/hasher"
	"ProjectDelivery/component/tokenprovider/jwt"
	"ProjectDelivery/modules/user/userbiz"
	"ProjectDelivery/modules/user/usermodel"
	"ProjectDelivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var dataLogin usermodel.UserLogin

		if err := c.ShouldBind(&dataLogin); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		hasher:= hasher.NewMd5Hash()
		//biz :=  userbiz.NewLoginBusiness(store,tokenProvider,hasher,30 )  ta sẽ ko set cứng thời gian hết hạn ..mà để cho devOps quyết định
		biz :=  userbiz.NewLoginBusiness(store,tokenProvider,hasher, 60 *60 *24 * 7)

		result ,err := biz.Login(c.Request.Context() , &dataLogin)

		if err != nil{
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}