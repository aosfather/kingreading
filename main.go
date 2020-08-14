package main

import (
	"github.com/aosfather/bingo_mvc/context"
	"github.com/aosfather/bingo_mvc/fasthttp"
	"github.com/aosfather/kingreading/apis"
	"github.com/aosfather/kingreading/collect"
	"github.com/aosfather/kingreading/profiles"
	"github.com/aosfather/kingreading/pushes"
)

func main() {
	dispatch := &fasthttp.FastHTTPDispatcher{}
	boot := context.Boot{}
	boot.Init(dispatch, load)
	boot.Start()
}

func load() []interface{} {

	return []interface{}{&collect.FileCollectManager{}, &pushes.PusherManager{}, &profiles.FileProfilesManager{}, &pushes.KindlePusher{}, &apis.Api{}, &collect.WuxianSpider{}}
}

//func (this *Application) Onload(context *bingo.ApplicationContext) bool {
//
//	//构建采集
//	f := collect.FileCollectManager{}
//	f.Init(context)
//
//	context.RegisterService("collects", &f)
//	//增加收集者
//	sp := &collect.WuxianSpider{}
//	f.AddCollecter("wuxianxs", sp)
//
//	//加载所有工作的主题
//	f.Load()
//
//	//采集定时任务
//	job := core.Job{}
//	job.Name = "collects"
//	job.Cron = "*/60 * * * ?"
//	job.Handler = f.JobHandle
//
//	jm := core.JobManager{}
//	context.RegisterService("job", &jm)
//	jm.Add(&job)
//
//	//构建push manager和kindle push
//	pm := pushes.PusherManager{}
//	//profiles manager
//	profilesMan := profiles.FileProfilesManager{}
//	profilesMan.Init(context)
//
//	context.RegisterService("profiles", &profilesMan)
//	context.RegisterService("pushers", &pm)
//
//	pm.Init(context)
//
//	p := pushes.KindlePusher{}
//	p.Load(context)
//
//	//推送定时任务
//	pjob := core.Job{}
//	pjob.Name = "kindlepusher"
//	pjob.Cron = "*/1 * * * ?"
//	pjob.Handler = pm.PushCronHandler
//	jm.Add(&pjob)
//
//	jm.Start(120)
//
//	log.Println("load fininshed")
//	return true
//
//}
