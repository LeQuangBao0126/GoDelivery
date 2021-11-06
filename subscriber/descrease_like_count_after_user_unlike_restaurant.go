package subscriber

import (
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"ProjectDelivery/pubsub"
	"context"
)

func RunDescreaseLikeCountAfterUserUnLikeRestaurant(appCtx component.AppContext) consumerJob{
	store:= restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
	return  consumerJob{
		Title: "Desc Like Count After User UnLike Restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			if likeData , ok := message.Data().(HasRestaurantId); ok {
				return store.DecreaseLikeCount(ctx,likeData.GetRestaurantId())
			}
			return nil
		},
	}
}