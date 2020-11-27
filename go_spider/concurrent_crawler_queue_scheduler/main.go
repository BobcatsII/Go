package main

import (
	"go_spider/concurrent_crawler_queue_scheduler/engine"
	"go_spider/concurrent_crawler_queue_scheduler/persist"
	"go_spider/concurrent_crawler_queue_scheduler/scheduler"
	"go_spider/concurrent_crawler_queue_scheduler/zhenai/parser"
)

func main() {

	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker, //单机版直接调worker
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, "ParseCityList"),
	})
}