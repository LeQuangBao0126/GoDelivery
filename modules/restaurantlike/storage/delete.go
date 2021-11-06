package restaurantlikestorage

import (
	"ProjectDelivery/common"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context , userid int ,restaurantid int ) error{
	db := s.db
	if err := db.Table("restaurant_likes").
	    Where("restaurant_id = ? and user_id= ?",restaurantid , userid).
		Delete(nil).Error  ; err!= nil{
		return common.ErrDB(err)
	}
	return nil
}