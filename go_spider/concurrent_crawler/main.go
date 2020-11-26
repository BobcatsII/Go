package main

import (
	"go_spider/concurrent_crawler/engine"
	"go_spider/concurrent_crawler/scheduler"
	"go_spider/concurrent_crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,    //10个worker线程
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
