package collect

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	"strings"
)

//无限小说抓取
type WuxianSpider struct {
	AbstractSpider
}

func (this *WuxianSpider) Run() {
	this.AbstractSpider.Run(this.GrabCaption)
}

//抓取
func (this *WuxianSpider) GrabCaption(c *Caption) {
	if c == nil || c.Enabled == false {
		return
	}

	if c.Url == "" {
		//获取书籍的地址
		c.Url = this.getCaptionUrl(c.Name)
	}

	if c.Url == "" {
		log.Println("not found book ", c.Name)
		return
	}

	//获取caption 对应的章节list
	captions := this.grabIndex(c.Url)
	currentIndex := len(captions) //最新章节的index
	fmt.Println(captions)
	fmt.Println(currentIndex)

	//检查caption的index,是否已经更新
	if c.Index <= 0 || c.Index < (currentIndex+1) {
		index := c.Index

		for index <= currentIndex {
			index++
			indexUrl := captions[c.Index-1]
			//抓取对应章节内容，更新index
			if !strings.HasPrefix(indexUrl, "http://") {
				indexUrl = c.Url + "/" + indexUrl
			}

			if this.grabContent(this.CaptionPath+"/"+c.Name, index, indexUrl) {
				c.Index = index
				c.IndexUrl = indexUrl
			} else {
				return
			}
		}

	}

}

//获取书籍对应的url
func (this *WuxianSpider) getCaptionUrl(name string) string {

	return ""
}

//返回章节列表
func (this *WuxianSpider) grabIndex(url string) []string {
	doc := this.GetDocument(url)
	se := doc.Find("dl dd")
	//
	//index:=se.First()
	var urls []string
	var href string
	var exists bool
	se.Each(func(index int, s *goquery.Selection) {
		href, exists = s.Children().Attr("href")
		if exists {
			urls = append(urls, href)
		}

	})

	return urls //总数

}

//抓取内容
func (this *WuxianSpider) grabContent(p string, index int, url string) bool {
	doc2 := this.GetDocument(url)
	if doc2 == nil {
		return false
	}

	file, e := os.Create(fmt.Sprintf("%s/%d.txt", p, index))
	if e != nil {
		log.Println("create file error!")
		return false
	}
	//标题
	io.WriteString(file, doc2.Find("div.bookname h1").Text())
	io.WriteString(file, doc2.Find("div#content").Text())
	defer file.Close()
	return true
}
