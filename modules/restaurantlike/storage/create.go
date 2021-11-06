package restaurantlikestorage

import (
	"ProjectDelivery/common"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context , data *restaurantlikemodel.Like ) error{
	db := s.db
	if err := db.Table(data.TableName()).Create(data).Error ; err!= nil{
		return common.ErrDB(err)
	}
	return nil
}