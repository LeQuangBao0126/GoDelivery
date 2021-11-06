package ginuser

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/component/hasher"
	"ProjectDelivery/modules/user/userbiz"
	"ProjectDelivery/modules/user/usermodel"
	"ProjectDelivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)
        //chỉ nên trả về đúng userid thôi . đung trả thêm chi phí ..
		//nếu mún tlay them cái khác lien quan user thì  store khác
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}