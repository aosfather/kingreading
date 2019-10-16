package collect

import (
	"fmt"
	"github.com/aosfather/bingo"
	"io/ioutil"
	"log"
)

type abstractCollectManager struct {
	root       string
	collecters map[string]Collecter
}

//注册收集器
func (this *abstractCollectManager) AddCollecter(key string, c Collecter) {
	if key == "" || c == nil {
		return
	}

	if this.collecters == nil {
		this.collecters = make(map[string]Collecter)
	}

	c.SetPath(this.root)
	this.collecters[key] = c
}

//增加收集的主题
func (this *abstractCollectManager) AddCaption(c *Caption) {
	if c == nil {
		return
	}

	collecter := this.collecters[c.Collect]

	if collecter != nil {
		collecter.AddCaption(c)
	}

}

//定时收集任务处理
func (this *abstractCollectManager) JobHandle() {
	log.Println("job be called")
	if this.collecters != nil {
		log.Println("start job")
		//执行收集任务
		for _, c := range this.collecters {
			c.Run()
		}

	}

	log.Println("job finished")
}

type FileCollectManager struct {
	abstractCollectManager
}

func (this *FileCollectManager) Init(config *bingo.ApplicationContext) {
	this.root = config.GetPropertyFromConfig("collect.workpath")
	fmt.Println(this.root)
}

func (this *FileCollectManager) Load() {
	if this.root == "" {
		return
	}

	rd, err := ioutil.ReadDir(this.root)
	if err != nil {

		return
	}
	for _, fi := range rd {
		if fi.IsDir() {

		} else {
			filename := this.root + "/" + fi.Name()
			c := Caption{}
			c.Load(filename)
			this.AddCaption(&c)
		}
	}

}

func (this *FileCollectManager) GetCaption(name string) *Caption {
	for _, c := range this.collecters {
		caption := c.GetCaption(name)
		if caption != nil {
			return caption
		}
	}

	return nil
}

//获取对应索引的文件名称
func (this *FileCollectManager) GetFileName(name string, index int) string {
	c := this.GetCaption(name)
	if c != nil {
		return fmt.Sprintf("%s/%s/%d.%s", this.root, c.Name, index, c.Fix)
	}
	return ""
}
