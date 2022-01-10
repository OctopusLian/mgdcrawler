/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-10 17:53:03
 */
package main

import (
	"net/http"

	"mgdcrawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/home/neo/Code/go/src/github.com/OctopusLian/mgdcrawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("/home/neo/Code/go/src/github.com/OctopusLian/mgdcrawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
