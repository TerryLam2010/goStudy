package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main(){
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	/*e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc : parser.ParseCitylist,
	})*/

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc : parser.ParseCity,
	})
}
