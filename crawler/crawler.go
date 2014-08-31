// Copyright (C) 2014 Frank Guan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fguan/ginger-collections"
)

func GetNextTarget(page string) (url string, endPos int) {
	url = ""
	endQuote := 0
	if startLink := strings.Index(page, "<a href="); startLink >= 0 {
		startQuote := strings.Index(page[startLink:], "\"") + startLink
		endQuote = strings.Index(page[startQuote+1:], "\"") + startQuote + 1
		//fmt.Println("==================")
		//fmt.Printf("%d %d\n",startQuote, endQuote)
		url = page[startQuote+1:endQuote]
		//fmt.Printf("%s\n", url)
	}
	return url, endQuote
}

func GetAllLinks(page string) []string {
	links :=  []string{}
	for {
		url, endPos := GetNextTarget(page)
		if url != "" {
			//fmt.Println(url)
			links = append(links, url)
			page = page[endPos:]
		} else {
			break
		}
	}
	return links
}

func GetPage(url string) (page string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(p)
}

func Crawl(seed string) map[string][]string {
	toCrawl := stack.New()
	toCrawl.Push(seed)
	crawled := map[string]bool{}
	index := map[string][]string{}

	for ; !toCrawl.Empty(); {
		page := toCrawl.Pop().(string)
		if crawled[page] == false {
			content = GetPage(page)
			AddPageToIndex(index, page, content)
			for _, v := range GetAllLinks(GetPage(page)) {
				toCrawl.Push(v)
			}
			crawled[page] = true
		}
	}
	return index
}
