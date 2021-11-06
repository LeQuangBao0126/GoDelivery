package restaurantbiz

import (
	"ProjectDelivery/common"
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"context"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context ,
		condition map[string] interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string ,
	) ([]restaurantmodel.Restaurant , error)
}
type LikeStore interface {
	GetRestaurantLike(ctx context.Context,ids []int) (map[int]int ,error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantBiz (store ListRestaurantStore,likeStore LikeStore ) *listRestaurantBiz {
	return &listRestaurantBiz{ store : store,likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context ,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	) ([]restaurantmodel.Restaurant , error) {

	result ,err := biz.store.ListDataByCondition(ctx , nil,filter,paging,"User")

	if err != nil{
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName,err)
	}
	////dung append lau hon
	//ids := make([]int , len(result))
	//
	//for i := range result{
	//	ids[i] = result[i].Id
	//}
	//mapResLike , err := biz.likeStore.GetRestaurantLike(ctx , ids)
	//
	//if err != nil{
	//	fmt.Println("Khong lay dc luot like nhưng vẫn trả ra List Restaurant")
	//}
	//
	//if v := mapResLike ; v != nil{
	//	for i,item := range result {
	//		result[i].LikedCount = mapResLike[item.Id]
	//	}
	//}

	return result, nil
}