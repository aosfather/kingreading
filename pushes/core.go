package pushes

import "log"

//pusher 接口定义
type Pusher interface {
	//执行推送
	//推送内容标题，文件列表
	Execute(config map[string]string, caption string, filelist ...string)
}

type PusherManager struct {
	pushers map[string]Pusher
}

//注册pusher处理器
func (this *PusherManager) Add(name string, p Pusher) {
	if name == "" || p == nil {
		log.Println("pusher is nil!", name)
		return
	}
	if this.pushers == nil {
		this.pushers = make(map[string]Pusher)
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
