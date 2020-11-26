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
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//创建worker
	in := make(chan Request) //公用输入\输出
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out) //worker就create好了
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
