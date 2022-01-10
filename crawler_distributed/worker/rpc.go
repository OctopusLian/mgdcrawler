/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 19:03:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 19:14:05
 */
package worker

import "mgdcrawler/engine"

type CrawlService struct{}

func (CrawlService) Process(
	req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
