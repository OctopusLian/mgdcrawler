/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 19:07:54
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 19:07:54
 */
package client

import (
	"net/rpc"

	"mgdcrawler/crawler_distributed/config"
	"mgdcrawler/crawler_distributed/worker"
	"mgdcrawler/engine"
)

func CreateProcessor(
	clientChan chan *rpc.Client) engine.Processor {

	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc,
			sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult),
			nil
	}
}
