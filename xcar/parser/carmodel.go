/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:50:11
 */
package parser

import (
	"regexp"

	"mgdcrawler/config"
	"mgdcrawler/engine"
)

var carDetailRe = regexp.MustCompile(`<a href="(/m\d+/)" target="_blank"`)

func ParseCarModel(
	contents []byte, _ string) engine.ParseResult {
	matches := carDetailRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: host + string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCarDetail, config.ParseCarDetail),
			})
	}

	return result
}
