package restaurantbiz

import (
	"ProjectDelivery/modules/restaurant/restaurantmodel"
	"context"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context ,data * restaurantmodel.RestaurantCreate) error
}

type restaurantBiz struct {
	store  CreateRestaurantStore
}

func NewCreateRestaurantBiz (store CreateRestaurantStore) *restaurantBiz{
	return &restaurantBiz{ store : store }
}

func (res *restaurantBiz) CreateRestaurant(ctx context.Context , data *restaurantmodel.RestaurantCreate) error{
	if err := res.store.Create(ctx , data) ; err != nil{
		return err
	}
	return nil
}