package restaurantlikestorage

import (
	"ProjectDelivery/common"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	"context"
)
//cai nay la lay map[restaurant] so luot like
func (s *sqlStore) GetRestaurantLike(ctx context.Context,ids []int) (map[int]int ,error){
	result := make(map[int]int )

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikedCount int `gorm:"column:count"`
	}
	var listLikeRestaurant []sqlData

	if err:= s.db.Table("restaurant_likes").Select("restaurant_id as restaurant_id , count(restaurant_id) as count").
	     Where("restaurant_id in (?)",ids).
		 Group("restaurant_id").
		 Find(&listLikeRestaurant).Error; err!= nil{

		 return nil,common.ErrDB(err)
	}

	for _ ,item:=  range listLikeRestaurant {
		result[item.RestaurantId] =  item.LikedCount
	}
	return result , nil
}

//di lay user likenhahang
func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context ,
	condition map[string] interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string ,) ( []common.SimpleUser ,error){

	var result []restaurantlikemodel.Like

	db := s.db.Table("restaurant_likes").Where(condition)



	if filter!= nil{
		if filter.RestaurantId >0 {
			db = db.Where("restaurant_id = ?" , filter.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Joins("User")



	if err:= db.Find(&result).Error ; err!= nil{
		return nil, common.ErrDB(err)
	}
	var users []common.SimpleUser
	for i , _ := range result{
		if result[i].User != nil{
			users = append(users, *result[i].User )
		}
	}
	return users , nil
}
