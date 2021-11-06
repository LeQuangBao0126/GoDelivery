package subscriber

import (
	"ProjectDelivery/component"
	"context"
)

func Setup(ctx component.AppContext){
	IncreaseLikeCountAfterUserLikeRestaurant(ctx,context.Background())
}