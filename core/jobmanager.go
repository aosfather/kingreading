package core

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

//任务handler 方法
type JobHandler func()

//定时任务管理
type Job struct {
	Name     string
	Cron     string
	Handler  JobHandler
	executer *cron.Cron
}

func (this *Job) start() {
	if this.executer == nil {
		this.executer = cron.New()
	}

	this.executer.AddFunc(this.Cron, func() {
		now := time.Now()
		if this.Handler != nil {
			this.Handler()
		}
		log.Println("cron running job:", this.Name, now.Minute(), now.Second())
	})
	this.executer.Start()

}

func (this *Job) stop() {
	if this.executer != nil {
		this.executer.Stop()
		now := time.Now()
		log.Println("cron stop job:", this.Name, now.Minute(), now.Second())
	}
}

type JobManager struct {
	jobs map[string]*Job
}

func (this *JobManager) Add(job *Job) {
	if job != nil {
		if this.jobs == nil {
			this.jobs = make(map[string]*Job)
		}
		this.jobs[job.Name] = job
		job.start()
	}
}

func (this *JobManager) Stop(jobname string) {
	if this.jobs != nil && jobname != "" {
		job := this.jobs[jobname]
		if job != nil {
			job.stop()
		}
	}
}

func (this *JobManager) Start(jobname string) {
	if this.jobs != nil && jobname != "" {
		job := this.jobs[jobname]
		if job != nil {
			job.start()
		}
	}
}
