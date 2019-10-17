package core

import "time"

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
	executer *JobExecuter
}

func (this *JobManager) Add(job *Job) {
	if job != nil {
		if this.jobs == nil {
			this.jobs = make(map[string]*Job)
		}
		this.jobs[job.Name] = job

		if this.executer == nil {
			this.executer = &JobExecuter{}
		}

		this.executer.AddJob(job)

	}
}

func (this *JobManager) Stop() {
	if this.executer != nil {
		this.executer.Stop()
	}
}

func (this *JobManager) Start(m int64) {
	if this.executer != nil {
		this.executer.Start(time.Duration(m) * time.Minute)
	}
}
