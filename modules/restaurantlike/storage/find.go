package restaurantlikestorage

import (
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	"context"
)

func (s *sqlStore) Find(ctx context.Context , userid int ,restaurantId int ) bool {
	db := s.db
 	var result restaurantlikemodel.Like
	if err := db.Table("restaurant_likes").
		Where("restaurant_id = ? and user_id= ?",restaurantId , userid).
		Find(&result).Error  ; err!= nil{
		 return false
	}
	if result.RestaurantId == 0 && result.UserId == 0 {
		return false
	}
	return true
}