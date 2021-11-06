package middleware

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"github.com/gin-gonic/gin"
)

func Recover (ac component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {
			defer func(){
				if err := recover() ; err!= nil{
					c.Header("Content-Type" , "application/json")
					if appErr ,ok := err.(*common.AppError); ok {
						c.AbortWithStatusJSON(appErr.StatusCode,appErr)
						 panic(err)//panic nay kich hoat cai stack trace error của ginDefault
						 return
					}

					appErr := common.ErrInternal(err.(error))
					c.AbortWithStatusJSON(appErr.StatusCode,appErr)
					panic(err) //panic nay kich hoat cai stack trace error của ginDefault
					return
				}
			}()
			c.Next()
	}
}