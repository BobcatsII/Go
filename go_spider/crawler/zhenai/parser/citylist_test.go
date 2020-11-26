package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	//验证结果和已知结果对比
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/baicheng1",
		"http://www.zhenai.com/zhenghun/cangzhou",
	}
	expectedCities := []string{
		"City: 阿坝",
		"City: 白城",
		"City: 沧州",
	}
	//验证URL的个数、前三个URL的正确性
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s", i, url, result.Requests[i].Url)
		}
	}

	//验证城市名的个数和前三个名称的正确性
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"request; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s;but was %s", i, city, result.Items[i].(string))
		}
	}
}
