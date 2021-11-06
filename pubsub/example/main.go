package main

import (
	"ProjectDelivery/pubsub"
	"ProjectDelivery/pubsub/pblocal"
	"context"
	"fmt"
	"time"
)

func main(){
	var localPb pubsub.Pubsub = pblocal.NewLocalPubsub()

	var topic pubsub.Topic = "Order Created"

	sub1 , _ := localPb.Subscribe(context.Background(),topic)
	sub2 ,_ := localPb.Subscribe(context.Background(),topic)

	localPb.Publish(context.Background(),topic,pubsub.NewMessage(1))
	localPb.Publish(context.Background(),topic,pubsub.NewMessage(2))

	go func(){
		for{
			fmt.Println("Con1" , (<-sub1).Data() )
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func(){
		for{
			fmt.Println("Con2" , (<-sub2).Data() )
			time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second *5)
}



