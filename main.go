/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 12:46:37
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 22:02:40
 */
package main

import (
	"mgdcrawler/config"
	"mgdcrawler/engine"
	"mgdcrawler/persist"
	"mgdcrawler/scheduler"
	"mgdcrawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
