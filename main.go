package main

import (
	"ProjectDelivery/component"
	"ProjectDelivery/middleware"
	"ProjectDelivery/modules/restaurant/restauranttransport/restaurantgin"
	"ProjectDelivery/modules/restaurantlike/transport/ginrestaurantlike"
	"ProjectDelivery/modules/user/usertransport/ginuser"
	"ProjectDelivery/pubsub/pblocal"
	"ProjectDelivery/subscriber"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loi khi load file env")
	}

	USERNAME := os.Getenv("username")
	PASSWORD := os.Getenv("PASSWORD")
	DatabaseHost := os.Getenv("DATABASE_HOST")
	DatabasePort := os.Getenv("DATABASE_PORT")
	DatabaseName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, DatabaseHost, DatabasePort, DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Loi khi ket noi database ")
	}
	runService(db)

}
func runService(db *gorm.DB) {
	db = db.Debug()
	SecretKey := os.Getenv("DATABASE_NAME")
	pb := pblocal.NewLocalPubsub()
	appCtx := component.NewAppContext(db,SecretKey,pb)

	//subscriber.Setup(appCtx)
	if err:= subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatalln("Dkm loi roi ......")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(middleware.Recover(appCtx))


	v1:= r.Group("/v1")
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	restaurants:= v1.Group("/restaurants",middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("" , restaurantgin.CreateRestaurant(appCtx))
		restaurants.GET("" , restaurantgin.ListRestaurant(appCtx))
		restaurants.GET("/:id" , restaurantgin.GetRestaurantById(appCtx))
		restaurants.PUT("/:id" , restaurantgin.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id" , restaurantgin.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/users-liked" , ginrestaurantlike.GetUsersLikeRestaurant(appCtx))
		restaurants.POST("/:id/like" , ginrestaurantlike.PostUsersLikeRestaurant(appCtx))
		restaurants.POST("/:id/unlike" , ginrestaurantlike.UserUnLikeRestaurant(appCtx))
	}



	r.Run()
}
