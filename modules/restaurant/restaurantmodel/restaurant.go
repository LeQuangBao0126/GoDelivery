package restaurantmodel

import "ProjectDelivery/common"

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
	Status int  `json:"status" gorm:"column:status"`
}
func(Restaurant) TableName()string {
	return "restaurants"
}

type RestaurantCreate struct {
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}
func(RestaurantCreate) TableName()string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}
func(RestaurantUpdate) TableName()string {
	return "restaurants"
}