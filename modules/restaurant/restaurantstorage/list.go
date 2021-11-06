package restaurantstorage

import (
	"ProjectDelivery/common"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context ,
	    condition map[string] interface{},
		filter *restaurantmodel.Filter,
	    paging *common.Paging,
		moreKeys ...string ,
		) ([]restaurantmodel.Restaurant , error){
	var result []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	db = db.Where(condition)

	if filter!= nil{
		if filter.CityId > 0 {
			db = db.Where("city_id = ? ", filter.CityId)
		}
	}
	if err:= db.Count(&paging.Total).Error ; err!= nil{
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys{
		db = db.Preload(moreKeys[i])
	}

	if  paging.FakeCursor != 1 && paging.FakeCursor >0{
		db = db.Where("id < ? ", paging.FakeCursor)
	}else{
		db = db.Offset((1-paging.Page)* paging.Limit)
	}

	if err:= db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error ; err!= nil{
		return nil, common.ErrDB(err)
	}

	paging.NextCursor = result[len(result)-1].Id

	return result, nil
}