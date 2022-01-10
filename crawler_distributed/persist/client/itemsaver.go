/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 18:35:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 18:36:25
 */
package client

import (
	"log"

	"mgdcrawler/crawler_distributed/config"
	"mgdcrawler/crawler_distributed/rpcsupport"
	"mgdcrawler/engine"
)

func ItemSaver(
	host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err := client.Call(
				config.ItemSaverRpc,
				item, &result)

			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return out, nil
}
