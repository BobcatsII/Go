package worker

import "go_spider/concurrent_crawler_queue_scheduler/engine"

type CrawlService struct{}

func (CrawlService) Process(
	req Request, result *ParseResult) error {
	engineReq,err := DeserializeRequest(req)
	if err!=nil{
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err!=nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}