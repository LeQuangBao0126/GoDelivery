package subscriber

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	"ProjectDelivery/modules/restaurant/restaurantstorage"
	"ProjectDelivery/pubsub"
	"context"
)
type HasRestaurantId interface {
	GetRestaurantId() int
}
func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext,ctx context.Context){
	c,_ := appCtx.GetPubsub().Subscribe(ctx,common.TopicUserLikeRestaurant)
 	store:= restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
	go func() {
		defer common.AppRecover()
		 for{
			 msg := <-c
			 likeData := msg.Data().(HasRestaurantId)
			 store.IncreaseLikeCount(ctx,likeData.GetRestaurantId())
		 }
	}()
}
//uoc gi
func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob{
	store:= restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
	return  consumerJob{
		Title: "Increase Like Count After User Like Restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			if likeData , ok := message.Data().(HasRestaurantId); ok {
				return store.IncreaseLikeCount(ctx,likeData.GetRestaurantId())
			}
			return nil
		},
	}
}

