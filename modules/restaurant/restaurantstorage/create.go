package restaurantstorage

import (
	"ProjectDelivery/common"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context ,data * restaurantmodel.RestaurantCreate) error{
	db := s.db
	if err := db.Table(data.TableName()).Create(data).Error ; err!= nil{
		return common.ErrDB(err)
	}
	return nil
}