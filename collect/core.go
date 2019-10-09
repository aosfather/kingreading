package collect

import (
	"encoding/json"
	"errors"
	"github.com/axgle/mahonia"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
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
func (this *Caption) Save(path string) {
	data, err := json.Marshal(this)
	if err == nil {
		err = ioutil.WriteFile(path+"/"+this.Name+".rc", data, 0644)
	}

}

//从文件加载数据
func (this *Caption) Load(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	} else {
		json.Unmarshal(data, this)
	}
}

type Collecter interface {
	AddCaption(c *Caption) //增加抓取主题
	Run()                  //执行抓取
}

//标准的爬虫
type AbstractSpider struct {
	CaptionPath string //配置文件的路径
	WorkPath    string //抓取文件存储路径
	captions    []*Caption
	captionsMap map[string]*Caption
}

func (this *AbstractSpider) Load() {
	if this.CaptionPath == "" {
		return
	}

	rd, err := ioutil.ReadDir(this.CaptionPath)
	if err != nil {

		return
	}
	for _, fi := range rd {
		if fi.IsDir() {

		} else {
			filename := this.CaptionPath + "/" + fi.Name()
			c := Caption{}
			c.Load(filename)
			this.AddCaption(&c)
		}
	}

}

//移除抓取的主题（并不删除文件，只是不再此队列中)
func (this *AbstractSpider) DelCaption(name string) {
	if name != "" {
		c := this.captionsMap[name]
		if c != nil {
			delete(this.captionsMap, name)
			for index, c1 := range this.captions {
				if c1 == c {
					this.captions = append(this.captions[:index], this.captions[index+1:]...)
					break
				}
			}
		}
	}
}

//增加抓取主题(加入队列中)
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
			caption.Save(this.CaptionPath)
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
