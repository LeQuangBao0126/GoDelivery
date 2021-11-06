package main

import (
	"fmt"
	"time"
)

func startConsumer(queue chan int , name string){
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(name,<-queue)
		}
	}()
}

func main(){
	queue:= make(chan int , 10000)

	for i := 0; i < 10000; i++ {
		queue<- i
	}
	startConsumer(queue,"Worker 1")
	startConsumer(queue,"Worker 2")
	startConsumer(queue,"Worker 3")

	time.Sleep(time.Second * 7)

}



