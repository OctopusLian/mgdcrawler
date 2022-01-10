/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 19:08:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 19:08:53
 */
package main

import (
	"fmt"

	"log"

	"flag"

	"mgdcrawler/crawler_distributed/rpcsupport"
	"mgdcrawler/crawler_distributed/worker"
	"mgdcrawler/fetcher"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
