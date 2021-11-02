package restaurantbiz

import (
	"context"
)

type DeleteRestaurantStore interface {

	Delete(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store  DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore ) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store :store }
}
func (biz *deleteRestaurantBiz)  DeleteRestaurant(ctx context.Context , id int )error {

	if err := biz.store.Delete(ctx , id) ; err!= nil{
		return err
	}
	return nil
}
