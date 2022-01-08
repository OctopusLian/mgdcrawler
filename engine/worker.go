/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:47:57
 */
package engine

import (
	"log"

	"mgdcrawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
