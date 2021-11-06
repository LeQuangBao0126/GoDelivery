package restaurantmodel

import (
	"ProjectDelivery/common"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string          `json:"name" gorm:"column:name"`
	UserId          int             `json:"-" gorm:"column:user_id"`
	Addr            string          `json:"addr" gorm:"column:addr"`
	Status          int             `json:"status" gorm:"column:status"`
	User            *common.SimpleUser    `json:"user"`
	LikedCount      int             `json:"liked_count" gorm:"liked_count"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"addr" gorm:"column:addr"`
	UserId int    `json:"-" gorm:"column:user_id"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}

//func ( data *Restaurant) Mask(isAdminOwner bool){
//	data.GenUID(common.DbTypeRestaurant)
//}
//khong ráng mà join restaurant và restaurant like tại vì bảng tobự mà join thì sẽ lâu ..và khi tách microservice khó khăn vì bị join rồi

