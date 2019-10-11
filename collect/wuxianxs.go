package collect

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

//无限小说抓取
type WuxianSpider struct {
	AbstractSpider
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

	//检查caption的index,是否已经更新
	if c.Index <= 0 || c.Index < currentIndex {
		index := c.Index

		for index < currentIndex {
			index++
			indexUrl := captions[c.Index-1]
			//抓取对应章节内容，更新index
			if this.grabContent(indexUrl) {
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
	file, _ := os.Open(url)

	root, _ := html.Parse(file)
	doc := goquery.NewDocumentFromNode(root)
	se := doc.Find("dl dd")
	fmt.Println(len(se.Nodes))                      //总数
	fmt.Println(se.First().Children().Attr("href")) //获取连接地址
	fmt.Println(len(se.Nodes))
	fmt.Println(se.Text())

	return nil
}

//抓取内容
func (this *WuxianSpider) grabContent(url string) bool {
	file, _ := os.Open(url)

	root, _ := html.Parse(file)
	doc := goquery.NewDocumentFromNode(root)

	se := doc.Find("td.line-content")
	fmt.Println(se.Text())

	root, _ = html.Parse(strings.NewReader(se.Text()))
	doc2 := goquery.NewDocumentFromNode(root)
	//标题
	fmt.Println(doc2.Find("div.bookname h1").Text())
	//内容
	fmt.Println(doc2.Find("div#content").Text())
	return false
}
