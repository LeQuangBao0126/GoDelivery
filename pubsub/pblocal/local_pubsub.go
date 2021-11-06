package pblocal

import (
	"ProjectDelivery/common"
	"ProjectDelivery/pubsub"
	"context"
	"fmt"
	"sync"
)

type localPubsub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubsub() *localPubsub {
	pb := &localPubsub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message, 10000),
		locker:       new(sync.RWMutex),
	}
	pb.run()
	return pb
}

func (ps *localPubsub) Publish(ctx context.Context, channel pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(channel)

	go func(){
		defer common.AppRecover()
		ps.messageQueue <- data
		fmt.Println("new event published =>>>>>",data.String())
	}()
	return nil
}

func (ps *localPubsub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
 	 c := make(chan *pubsub.Message)
	 ps.locker.Lock()

	 if  val,ok := ps.mapChannel[topic] ;ok{
		 val = append(ps.mapChannel[topic] , c)
		 ps.mapChannel[topic] = val
	 }else{
		 ps.mapChannel[topic] = []chan *pubsub.Message{c}
	 }
	 ps.locker.Unlock()

	 return c , func() {
		 fmt.Println("UnSubscribe ")
		 if chanArr , ok := ps.mapChannel[topic] ; ok {
			 for i := range chanArr{
				 if chanArr[i] == c {
					 newChanArr:= append (chanArr[:i],chanArr[i+1 :]... )
					 ps.locker.Lock()
					 ps.mapChannel[topic] = newChanArr
					 ps.locker.Unlock()
					 break
				 }
			 }
		 }
	 }

}

func (ps *localPubsub) run() {
	fmt.Println("Pubsub start Run")
	go func() {
		for {
			mess := <-ps.messageQueue
			fmt.Println("Message broker is dequeue ", mess.Data())
			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					//neu o subcribe chua rut ra ..thi se ket ca he thong nen xai go
					go func(c chan *pubsub.Message) {
						c<- mess
					}(subs[i])
				}
			}
		}
	}()
}
