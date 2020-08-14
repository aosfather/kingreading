package apis

import (
	"fmt"
	"github.com/aosfather/bingo_mvc"
	"github.com/aosfather/kingreading/collect"
	"github.com/aosfather/kingreading/pushes"
)

/**
  发送内容
*/

type Request struct {
}
type Api struct {
	Collector   *collect.FileCollectManager `Inject:"" mapper:"name(collect);url(/fresh);method(POST);style(JSON)"`
	PushManager *pushes.PusherManager       `Inject:"" mapper:"name(push);url(/push);method(POST);style(JSON)"`
}

func (this *Api) GetHandles() bingo_mvc.HandleMap {
	result := bingo_mvc.NewHandleMap()
	result.Add("collect", this.collect, &Request{})
	result.Add("push", this.push, &Request{})
	return result
}

func (this *Api) collect(a interface{}) interface{} {
	r := a.(*Request)
	go this.Collector.JobHandle()
	fmt.Println(r)
	return "ok"
}

func (this *Api) push(a interface{}) interface{} {
	r := a.(*Request)
	this.PushManager.PushCronHandler()
	fmt.Println(r)
	return "ok"
}
