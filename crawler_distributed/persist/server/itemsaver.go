/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 18:35:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 18:37:33
 */
package main

import (
	"flag"
	"fmt"
	"log"

	"mgdcrawler/config"
	"mgdcrawler/crawler_distributed/persist"
	"mgdcrawler/crawler_distributed/rpcsupport"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
