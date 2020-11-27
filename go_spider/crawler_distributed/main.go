package main

import (
	"flag"
	"fmt"
	"go_spider/concurrent_crawler_queue_scheduler/engine"
	"go_spider/concurrent_crawler_queue_scheduler/persist"
	"go_spider/concurrent_crawler_queue_scheduler/scheduler"
	"go_spider/concurrent_crawler_queue_scheduler/zhenai/parser"
	"go_spider/crawler_distributed/config"
	itemsaver "go_spider/crawler_distributed/persist/client"
	"go_spider/crawler_distributed/rpcsupport"
	worker "go_spider/crawler_distributed/worker/client"
	"log"
	"net/rpc"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {

	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//Scheduler:&scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		//RequestProcessor: engine.Worker,
		RequestProcessor: processor,
	}
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	PaserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		//Url:       "http://www.zhenai.com/zhenghun/shanghai",
		Url: "http://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCity,
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})

	//flag.Parse()

	//concurrentScheduler() //并发队列版，每个worker单独用一个channel
	//simpleScheduler()  //并发非队列版，公用一个channel
	//singleCity()  //单个城市直接请求
}

//func concurrentScheduler() {
//	itemsChan, err := itemsaver.ItemSaver(*itemSaverHost)
//	if err != nil {
//		panic(err)
//	}
//
//	pool := createClientPool(strings.Split(*workerHosts, ", "))
//	processor := worker.CreateProcessor(pool)
//
//	e := engine.ConcurrentEngine{
//		Scheduler:        &scheduler.QueuedScheduler{},
//		WorkerCount:      10,
//		ItemChan:         itemsChan,
//		RequestProcessor: processor,
//	}
//
//	e.Run(engine.Request{
//		Url:    "http://www.zhenai.com/zhenghun",
//		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
//	})
//}

func simpleScheduler() {
	itemsChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan:    itemsChan,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func singleCity() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})

}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s：%v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
