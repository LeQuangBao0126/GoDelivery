package subscriber

import (
	"ProjectDelivery/common"
	"ProjectDelivery/component"
	asyncjob "ProjectDelivery/component/async_job"
	"ProjectDelivery/pubsub"
	"context"
	"fmt"
)

type consumerJob struct {
	Title string
	Hld func(ctx context.Context, message *pubsub.Message) (error)
}

type consumerEngine struct {
	appCtx component.AppContext
}

func NewEngine (appContext component.AppContext) *consumerEngine{
	return &consumerEngine{appCtx: appContext}
}
func (engine *consumerEngine) Start()error {

	engine.startSubTopic(common.TopicUserLikeRestaurant,
					true,
					RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx))
	engine.startSubTopic(common.TopicUserUnLikeRestaurant,
		true,
		 RunDescreaseLikeCountAfterUserUnLikeRestaurant(engine.appCtx))
	return nil
}

type GroupJob interface {
	Run(ctx context.Context)error
}
func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _,item := range consumerJobs{
		fmt.Println("Setup consumer for : ",item.Title)
	}
	getJobHandler := func(job *consumerJob,message *pubsub.Message) asyncjob.JobHandler{
		return func(ctx context.Context) error {
			fmt.Println("running job for ",job.Title , "value ",message.Data())
			return job.Hld(ctx,message)
		}
	}
	go func(){
		for{
			msg := <-c
			jobHdlArr:= make([]asyncjob.Job , len(consumerJobs))

			for i:= range consumerJobs{
				jobHdlArr[i] = asyncjob.NewJob(getJobHandler(&consumerJobs[i],msg))
			}

			group:= asyncjob.NewGroup(isConcurrent,jobHdlArr...)
			if err:= group.Run(context.Background()); err!= nil {
				fmt.Println(err)
			}
		}
	}()
	return nil
}