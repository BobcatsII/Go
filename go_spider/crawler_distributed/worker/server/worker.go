package main

import (
	"fmt"
	"go_spider/crawler_distributed/config"
	"go_spider/crawler_distributed/rpcsupport"
	"go_spider/crawler_distributed/worker"
	"log"
)

//命令行库，可以读取加入的参数
// go run worker.go --port=9000
//var port = flag.Int("port",0,"the port for me to listen on")
func main() {
	//flag.Parse()
	//if *port == 0{
	//	fmt.Println("must specify a port")
	//}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d",config.WorkerPort0),
		worker.CrawlService{}))
	//log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),worker.CrawlService{}))
}


