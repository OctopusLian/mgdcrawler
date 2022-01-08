/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 12:46:37
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 12:57:44
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	//乱码转中文
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s\n", string(all))
}
