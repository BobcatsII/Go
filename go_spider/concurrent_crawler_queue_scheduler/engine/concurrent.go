package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler //scheduler从这里来
	WorkerCount int       //创建多少个worker进行统计
}

//scheduler是个接口
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//创建worker
	out := make(chan ParseResult) //只要一个out就够了
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out,e.Scheduler) //worker就create好了
	}

	//将seeds全部传给scheduler
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	//收out，out是worker输出的结果
	for {
		result := <-out //从out里面收进来
		//打印Items
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		//打印完Item之后，把result里所有的request送给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			//tell scheduler i'm ready
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
