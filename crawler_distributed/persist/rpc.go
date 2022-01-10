/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-10 18:31:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 21:04:10
 */
package persist

import (
	"log"
	"mgdcrawler/engine"
	"mgdcrawler/persist"

	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item,
	result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return err
}
