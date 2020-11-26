package scheduler

import "go_spider/concurrent_crawler_queue_scheduler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	//每一个request会创建一个goroutine，作用是往统一的channle去分发，也就是workerChan
	go func() {s.workerChan <- r}()
}