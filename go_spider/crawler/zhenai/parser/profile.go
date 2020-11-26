package parser

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"go_spider/crawler/engine"
	"go_spider/crawler/model"
	"log"
	"regexp"
)

//json和html方法二选一：
//json过滤方法
const jsonRe = `<script>window.__INITIAL_STATE__=(.+);\(function`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(jsonRe)
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		profile := parseJson(json)
		profile.Name = name
		result.Items = append(result.Items, profile)
		fmt.Println(result)
	}
	return result
}

func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("parser json fail ...")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()

	var profile model.Profile
	for k, v := range infos {
		fmt.Printf("k:%v,%T\n", k, k)
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}
	}
	return profile
}

//////////////////////////////////////////////////////////
// HTML正则方法，写在这里方便回顾，重点是方法
//var ageRe = regexp.MustCompile(
//	`<td><span class="label">年龄: </span>([^>]+)</td>`)
//var marriageRe = regexp.MustCompile(
//	`<td><span class="label">婚况: </span>([^>]+)</td>`)
//... 根据model.profile的结构体，按需要添加需要获取的信息名称 ...
//
//func ParseProfile(contents []byte, name string) engine.ParseResult  {
//	profile := model.Profile{}
//	profile.Age = extractString(contents, ageRe)
//	profile.Marriage = extractString(contents, marriageRe)
//	... 上面的变量添加对应信息的代码 ...
//	//发现缺少了profile.Name字段
//	profile.Name = name
//
//	result := engine.ParseResult{
//		Items: []interface{}{profile},
//	}
//	return result
//}
//
//func extractString(contents []byte, re *regexp.Regexp) string {
//	match := re.FindSubmatch(contents)
//	if len(match) >= 2 {
//		return string(match[1])
//	} else {
//		return ""
//	}
//}
