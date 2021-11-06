package rstrtlikebiz

import (
	"ProjectDelivery/common"
	restaurantlikemodel "ProjectDelivery/modules/restaurantlike/model"
	"ProjectDelivery/pubsub"
	"context"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	Find(ctx context.Context , userid int ,restaurantId int ) bool
}

//type IncreaseLikeCountStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	pubsub    pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store,pubsub: pubsub }
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	if isExist := biz.store.Find(ctx, data.UserId , data.RestaurantId) ; isExist == true{
		return restaurantlikemodel.ErrorHasBeenExisted(nil)
	}
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	//từ module restaurant like ..dùng cái store này để giao tiếp vs restaurant để cap nhat liked_count
	//side effect . new soliution
	biz.pubsub.Publish(ctx,common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	//go func(){
	//	defer common.AppRecover()
	//	job1 := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.incStore.IncreaseLikeCount(ctx,data.RestaurantId)
	//	})
	//	asyncjob.NewGroup(true , job1).Run(ctx)
	//}()



	return nil
}