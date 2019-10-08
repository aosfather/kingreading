package collect

import (
	"errors"
	"github.com/axgle/mahonia"
	"golang.org/x/net/html"
	"net/http"
)

//标题记录
type Caption struct {
	Name    string //标题名称
	Title   string //书籍标题
	Index   int    //当前的索引号
	Enabled bool   //是否启用
	//
	Url      string //对应的url
	IndexUrl string //索引对应的地址
}

//保存
func (this *Caption) Save() {

}

type Collecter interface {
	AddCaption(c *Caption) //增加抓取主题
	Run()                  //执行抓取
}

//标准的爬虫
type AbstractSpider struct {
	captions    []*Caption
	captionsMap map[string]*Caption
}

//增加抓取主题
func (this *AbstractSpider) AddCaption(c *Caption) {
	if c != nil {
		if this.captionsMap == nil {
			this.captionsMap = make(map[string]*Caption)
		}

		if this.captionsMap[c.Name] == nil {
			this.captionsMap[c.Name] = c
			this.captions = append(this.captions, c)
		}

	}

}

//抓取
func (this *AbstractSpider) GrabCaption(c *Caption) {

}

//执行抓取
func (this *AbstractSpider) Run() {
	if len(this.captions) > 0 {
		for _, caption := range this.captions {
			this.GrabCaption(caption)
			caption.Save()
		}

	}

}

//获取可用的主题
func (this *AbstractSpider) GetCaption(name string) *Caption {
	if name != "" && this.captionsMap != nil {
		return this.captionsMap[name]
	}

	return nil
}

func (this *AbstractSpider) GetCaptionByIndex(index int) *Caption {
	if index < 0 || this.captions == nil || len(this.captions) == 0 || index >= len(this.captions) {
		return nil
	}

	return this.captions[index]
}

func (this *AbstractSpider) GetCaptionCount() int {
	return len(this.captions)
}

//获取url的内容
func (this *AbstractSpider) GetUrl(url string) (*html.Node, error) {
	res, e := http.Get(url)
	if e != nil {
		return nil, e
	}

	defer res.Body.Close()
	if res.Request == nil {
		return nil, errors.New("Response.Request is nil")
	}

	// Parse the HTML into nodes
	root, e := html.Parse(res.Body)
	if e != nil {
		return nil, e
	}
	return root, nil
}

//-------------------------字符编码转换 -----------------------------------//
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GbkToUtf8(src string) string {
	return ConvertToString(src, "gbk", "utf-8")
}
