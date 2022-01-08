/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 21:46:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 21:48:54
 */
package view

import (
	"html/template"
	"io"

	"mgdcrawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(
	filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(
	w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
