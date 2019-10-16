package pushes

import (
	"github.com/aosfather/bingo"
	"github.com/aosfather/kingreading/profiles"
	"log"
	"strings"
)

//pusher 接口定义
type Pusher interface {
	//执行推送
	//推送内容标题，文件列表
	Execute(p *profiles.Profile)
}

type PusherManager struct {
	catalogs []string
	p        profiles.ProfileManager
	pushers  map[string]Pusher
}

func (this *PusherManager) Init(context *bingo.ApplicationContext) {
	if this.pushers == nil {
		this.pushers = make(map[string]Pusher)
	}
	this.catalogs = strings.Split(context.GetPropertyFromConfig("profile.catalog"), ",")
	this.p = context.GetService("profiles").(profiles.ProfileManager)
}

//注册pusher处理器
func (this *PusherManager) Add(name string, p Pusher) {
	if name == "" || p == nil {
		log.Println("pusher is nil!", name)
		return
	}

	this.pushers[name] = p
}

//获取推送器
func (this *PusherManager) Get(rt string) Pusher {
	if rt == "" || this.pushers == nil {
		return nil
	}

	return this.pushers[rt]
}

//推送定时job的处理函数
func (this *PusherManager) PushCronHandler() {
	log.Println("start push")
	//循环profile,根据level，先处理level 低的。
	for _, catalog := range this.catalogs {
		count := this.p.GetProfileCount(catalog)
		log.Println("profile size=", count)
		for i := 0; i < count; i++ {
			p := this.p.GetProfile(catalog, i)
			if p != nil {
				pusher := this.Get(p.RemoteType)
				if pusher != nil {
					pusher.Execute(p)
				}
			}

		}

	}

}
