package parser

import (
	"go_spider/concurrent_crawler_queue_scheduler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if e != nil {
		panic(e)
	}

	result := ParseCityList(contents, "")
	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result sould have %d requests; but had %d",
			resultSize, len(result.Requests))
	}
}
