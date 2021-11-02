package restaurantstorage

import (
	"ProjectDelivery/common"
	"context"
)

func (s sqlStore) Delete(ctx context.Context, id int) error{
	db := s.db
	if err:=  db.Table("restaurants").Where("id = ? ",id).Delete(nil).Error ; err != nil{
		return common.ErrDB(err)
	}
	return nil
}