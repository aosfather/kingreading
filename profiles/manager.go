package profiles

import (
	"github.com/aosfather/bingo"
	"io/ioutil"
	"log"
	"os"
)

//profiles manager implements
type FileProfilesManager struct {
	path     string
	profiles map[string][]*Profile
}

func (this *FileProfilesManager) Init(context *bingo.ApplicationContext) {
	this.path = context.GetPropertyFromConfig("profile.workpath")
	this.profiles = make(map[string][]*Profile)
	this.load()
}

func (this *FileProfilesManager) load() {
	if this.path == "" {
		return
	}

	rd, err := ioutil.ReadDir(this.path)
	if err != nil {

		return
	}
	for _, fi := range rd {
		if fi.IsDir() {

		} else {
			filename := this.path + "/" + fi.Name()
			p := Profile{}
			f, e := os.Open(filename)
			if e != nil {
				log.Println(e)
			} else {
				p.Load(f)
				this.AddProfile(&p)
			}

		}
	}
}

func (this *FileProfilesManager) AddProfile(p *Profile) {
	if p == nil {
		return
	}
	catalog := p.Catalog
	p.ProfileMan = this
	list := this.profiles[catalog]
	list = append(list, p)
	this.profiles[catalog] = list

}

//持久化
func (this *FileProfilesManager) Save(p *Profile) {
	if p != nil && this.path != "" {
		f, e := os.Create(this.path + "/" + p.ID + ".pf")
		if e == nil {
			p.Save(f)
			f.Close()
		}
	}
}

//获取配置
func (this *FileProfilesManager) GetProfile(catalog string, index int) *Profile {
	list := this.profiles[catalog]
	if list != nil {
		if index >= 0 && index < len(list) {
			return list[index]
		}
	}
	return nil
}

//获取配置的个数
func (this *FileProfilesManager) GetProfileCount(catalog string) int {
	list := this.profiles[catalog]
	if list != nil {
		return len(list)
	}
	return 0
}
