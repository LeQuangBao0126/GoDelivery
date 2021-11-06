package rstrtlikebiz

import (
	"ProjectDelivery/common"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	"context"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context ,
		condition map[string] interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string ,) ( []common.SimpleUser ,error)
}
type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore ) *listUserLikeRestaurantBiz{
	return &listUserLikeRestaurantBiz{ store :store}
}

func(biz *listUserLikeRestaurantBiz)  ListUsers ( ctx context.Context ,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,) ([]common.SimpleUser ,error){

	result , err := biz.store.GetUsersLikeRestaurant(ctx,nil,filter,paging,"User")
	if err != nil{
		return nil , common.ErrCannotListEntity("Restaurant like",err)
	}
	return result ,nil
}
