package main

import (
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restauranttransport/restaurantgin"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	r := gin.Default()
	db = db.Debug()
	appCtx := component.NewAppContext(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	restaurants:= r.Group("/restaurants")
	{
		restaurants.POST("" , restaurantgin.CreateRestaurant(appCtx))
		restaurants.GET("" , restaurantgin.ListRestaurant(appCtx))
		restaurants.GET("/:id" , restaurantgin.GetRestaurantById(appCtx))
		restaurants.PUT("/:id" , restaurantgin.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id" , restaurantgin.DeleteRestaurant(appCtx))
	}
	r.Run()
}
