package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (s SimpleEngine)Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds {
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests= requests[1:]
		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,parserResult.Requestes...)
		for _,item := range parserResult.Items {
			log.Printf("Got item %v",item)
		}

	}
}

func worker(r Request)(ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Getcher:error getching url %s:%v",r.Url,err)
		return ParseResult{},err
	}
	return  r.ParserFunc(body),nil
}