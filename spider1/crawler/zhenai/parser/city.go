package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRex = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
 	urlRex = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
 )

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRex.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches {
		name := string(m[2])
		result.Items = append(
			result.Items,"User" + name)
		result.Requestes = append(
			result.Requestes,engine.Request{
				Url: string(m[1]),
				ParserFunc: func(contents []byte)engine.ParseResult{
					return ParseProfile(contents,name)
				},
			})
	}

	matches1 := urlRex.FindAllSubmatch(contents, -1)
	for _,m := range matches1 {
		result.Requestes = append(result.Requestes,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
