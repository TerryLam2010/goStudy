package engine

import (
	"crawler/model"
	"log"
)

type  ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i< e.WorkerCount;i++ {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}
	for _,r := range seeds {
		// 提交到  s.requestChan <- r 放入Scheduler 的chan 里面
		e.Scheduler.Submit(r)
	}
	profileCount := 0
	for {
		// 死循环等待chan 给数据
		result := <-out
		for _,item := range result.Items {
			if _,ok := item.(model.Profile);ok {
				log.Printf("Got item #%d %v",profileCount,item)
				profileCount++
			}

		}
		for _,r := range result.Requestes {
			if isDuplicate(r.Url){
				continue
			}
			// 把爬到的url 继续给到chan 继续爬
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier){
	go func() {
		for {
			ready.WorkerReady(in)
			// tell scheduler i'm ready
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}