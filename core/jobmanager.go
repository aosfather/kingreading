package core

import (
	"github.com/robfig/cron"
	"log"
)

//任务handler 方法
type JobHandler func()

//定时任务管理
type Job struct {
	Name    string
	Cron    string
	Handler JobHandler
}

type JobManager struct {
	jobs     map[string]*Job
	executer *cron.Cron
}

func (this *JobManager) Add(job *Job) {

	if job != nil {
		if this.jobs == nil {
			this.jobs = make(map[string]*Job)
		}
		this.jobs[job.Name] = job

		if this.executer == nil {
			this.executer = cron.New()

			this.executer.Start()
		}

		log.Println(this.executer.AddFunc(job.Cron, job.Handler))

	}
}

func (this *JobManager) Stop(jobname string) {
	//if this.jobs != nil && jobname != "" {
	//	job := this.jobs[jobname]
	//	if job != nil {
	//		this.executer.Remove(jobname)
	//	}
	//}
}

func (this *JobManager) Start(jobname string) {
	//if this.jobs != nil && jobname != "" {
	//	job := this.jobs[jobname]
	//	if job != nil {
	//		job.start()
	//	}
	//}
}
