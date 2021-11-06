package rstrtlikebiz

import (
	"ProjectDelivery/common"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	"ProjectDelivery/pubsub"
	"context"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userid int, restaurantId int) error
	Find(ctx context.Context, userid int, restaurantId int) bool
}
//type DecreaseLikeCountStore interface {
//	DecreaseLikeCount(ctx context.Context,id int) error
//}

type userUnLikeRestaurantBiz struct {
	store UserUnLikeRestaurantStore
	pubsub pubsub.Pubsub
	//descStore DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, pb pubsub.Pubsub  ) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store: store, pubsub:pb  }
}

func (biz *userUnLikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	if isExist := biz.store.Find(ctx, data.UserId, data.RestaurantId); isExist == false {
		return restaurantlikemodel.ErrorHasBeenUnlike(nil)
	}
	err := biz.store.Delete(ctx, data.UserId, data.RestaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	//side effect
	//go func() {
	//	defer common.AppRecover()
	//	biz.descStore.DecreaseLikeCount(ctx,data.RestaurantId)
	//
	//}()

	biz.pubsub.Publish(ctx,common.TopicUserUnLikeRestaurant, pubsub.NewMessage(data))

	return nil
}
