package main

import (
	"fmt"
	"github.com/aosfather/bingo"
	"github.com/aosfather/kingreading/collect"
	"github.com/aosfather/kingreading/core"
	"github.com/aosfather/kingreading/profiles"
	"github.com/aosfather/kingreading/pushes"
	"log"
)

func main() {
	fmt.Println("hello")
	app := Application{}
	app.Init()  //初始化
	app.start() //服务启动
}

type Application struct {
	bingo.Application
}

func (this *Application) Init() {
	this.SetHandler(this.Onload, nil)

}
func (this *Application) start() {
	this.RunApp()
}

func (this *Application) Onload(context *bingo.ApplicationContext) bool {

	//构建采集
	f := collect.FileCollectManager{}
	f.Init(context)

	context.RegisterService("collects", &f)
	//增加收集者
	sp := &collect.WuxianSpider{}
	f.AddCollecter("wuxianxs", sp)

	//加载所有工作的主题
	f.Load()

	//采集定时任务
	job := core.Job{}
	job.Name = "collects"
	job.Cron = "*/60 * * * ?"
	job.Handler = f.JobHandle

	jm := core.JobManager{}
	context.RegisterService("job", &jm)
	jm.Add(&job)

	//构建push manager和kindle push
	pm := pushes.PusherManager{}
	//profiles manager
	profilesMan := profiles.FileProfilesManager{}
	profilesMan.Init(context)

	context.RegisterService("profiles", &profilesMan)
	context.RegisterService("pushers", &pm)

	pm.Init(context)

	p := pushes.KindlePusher{}
	p.Load(context)

	//推送定时任务
	pjob := core.Job{}
	pjob.Name = "kindlepusher"
	pjob.Cron = "*/50 * * * ?"
	pjob.Handler = pm.PushCronHandler
	jm.Add(&pjob)

	//加载profile
	pone := profiles.Profile{}
	pone.Init()
	pone.Caption = "xs_wgsd"
	pone.RemoteType = "kindle"
	pone.LastSendIndex = 2437
	pone.MaxLimit = 2
	pone.Rate = profiles.RT_HOUR
	pone.LastHour = 10
	pone.LastMinutes = 56
	pone.ExtProperties["EMAIL"] = "faye.feelcool@kindle.cn"
	pone.Catalog = "0"
	pone.ID = "user001"
	profilesMan.AddProfile(&pone)

	log.Println("load fininshed")
	return true

}
