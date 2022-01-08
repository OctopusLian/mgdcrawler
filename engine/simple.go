/*
 * @Description: 调度器简单版本实现
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:58:04
 */

package engine

import (
	"log"
)

type SimpleEngine struct{} //空结构体作用_不消耗内存

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests,
			parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
