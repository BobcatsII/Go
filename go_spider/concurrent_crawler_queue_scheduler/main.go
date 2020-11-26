package main

import (
	"go_spider/concurrent_crawler_queue_scheduler/engine"
	"go_spider/concurrent_crawler_queue_scheduler/scheduler"
	"go_spider/concurrent_crawler_queue_scheduler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,    //10个worker线程
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
