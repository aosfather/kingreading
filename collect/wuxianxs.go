package collect

import "log"

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

	return nil
}

//抓取内容
func (this *WuxianSpider) grabContent(url string) bool {

	return false
}
