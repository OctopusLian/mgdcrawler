/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:50:43
 */
package parser

import (
	"regexp"

	"mgdcrawler/config"
	"mgdcrawler/engine"
)

const cityListRe = `<a href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(
	contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCity),
			})
	}

	return result
}
