package parser

import (
	"go_spider/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//解析城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	i := 0
	for _, m := range matches {
		result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//太多了，不方便查看效果，弄1个城市就行了
		i++
		if i == 1 {
			break
		}
	}
	return result
}
