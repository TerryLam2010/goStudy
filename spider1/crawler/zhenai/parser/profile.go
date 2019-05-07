package parser

import (
	"crawler/engine"
	"crawler/model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"strconv"
	"strings"
)


var JsonRex = regexp.MustCompile(`window.__INITIAL_STATE__=([^;]+);\(function\(\)`)
func ParseProfile(contents []byte,name string) engine.ParseResult{
	profile := model.Profile{}
	json_str := extractString(contents, JsonRex)
	if json_str == "" {
		return engine.ParseResult{}
	}
	json, _ := simplejson.NewJson([]byte(json_str))
	objectInfo := json.Get("objectInfo")
	if objectInfo == nil {
		log.Fatal(json_str)
	}
	age, e := objectInfo.Get("age").Int()
	if e == nil {
		profile.Age = age
	}
	heightStr,_ := objectInfo.Get("heightString").String()
	height, e := strconv.Atoi(strings.Replace(heightStr,"cm","",-1))
	if e == nil {
		profile.Height = height
	}
	profile.Weight = 120

	profile.Name = name
	profile.Income = objectInfo.Get("salaryString").MustString()
	profile.Gender = objectInfo.Get("genderString").MustString()
	basicInfo, _ := objectInfo.Get("basicInfo").StringArray()
	profile.Car = "无"
	profile.Xinzuo = basicInfo[2]
	profile.Education = objectInfo.Get("educationString").MustString()
	profile.Occupation = ""
	profile.Hokou = objectInfo.Get("workProvinceCityString").MustString()
	profile.House = "无"
	return engine.ParseResult{Items: []interface{}{profile}}
}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1  {
		return string(match[1])
	}else{
		return ""
	}
}