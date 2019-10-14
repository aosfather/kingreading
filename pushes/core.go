package pushes

import (
	"github.com/aosfather/kingreading/profiles"
	"log"
)

//pusher 接口定义
type Pusher interface {
	Config(config map[string]string)
	//执行推送
	//推送内容标题，文件列表
	Execute(p *profiles.Profile)
}

type PusherManager struct {
	catalogs []string
	p        profiles.ProfileManager
	pushers  map[string]Pusher
}

func (this *PusherManager) Init() {
	if this.pushers == nil {
		this.pushers = make(map[string]Pusher)
	}
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

var _pushers *PusherManager

//设置发送者
func SetPushers(p *PusherManager) {
	_pushers = p
}

//推送定时job的处理函数
func PushCronHandler() {
	if _pushers == nil {
		return
	}
	//循环profile,根据level，先处理level 低的。
	for _, catalog := range _pushers.catalogs {
		count := _pushers.p.GetProfileCount(catalog)

		for i := 0; i < count; i++ {
			p := _pushers.p.GetProfile(catalog, i)
			if p != nil {
				pusher := _pushers.Get(p.RemoteType)
				if pusher != nil {
					pusher.Execute(p)
				}
			}

		}

	}

}
