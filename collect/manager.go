package collect

import (
	"fmt"
	"io/ioutil"
	"log"
)

type FileCollectManager struct {
	Root       string `Value:"collect.workpath"`
	collecters map[string]Collecter
}

func (this *FileCollectManager) Init() {
}

//注册收集器
func (this *FileCollectManager) AddCollecter(key string, c Collecter) {
	if key == "" || c == nil {
		return
	}

	if this.collecters == nil {
		this.collecters = make(map[string]Collecter)
	}

	c.SetPath(this.Root)
	this.collecters[key] = c
}

//增加收集的主题
func (this *FileCollectManager) AddCaption(c *Caption) {
	if c == nil {
		return
	}

	collecter := this.collecters[c.Collect]

	if collecter != nil {
		collecter.AddCaption(c)
	}

}

//定时收集任务处理
func (this *FileCollectManager) JobHandle() {
	log.Println("collect job be called!")
	if this.collecters != nil {
		log.Println("start job")
		//执行收集任务
		for _, c := range this.collecters {
			c.Run()
		}

	}

	log.Println("job finished")
}

func (this *FileCollectManager) Load() {
	if this.Root == "" {
		return
	}

	rd, err := ioutil.ReadDir(this.Root)
	if err != nil {

		return
	}
	for _, fi := range rd {
		if fi.IsDir() {
			log.Println("ignore dir")
		} else {
			filename := this.Root + "/" + fi.Name()
			c := Caption{}
			c.Load(filename)
			log.Println(c)
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
		return fmt.Sprintf("%s/%s/%d.%s", this.Root, c.Name, index, c.Fix)
	}
	return ""
}
