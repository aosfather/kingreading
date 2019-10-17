package core

import (
	"fmt"
	"log"
	"time"
)

//job实现
type JobExecuter struct {
	jobs []*Job
	quit chan int
}

func (this *JobExecuter) Start(d time.Duration) {
	ticker := time.NewTicker(d)
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("run job")
				this.run()

			case <-quit:
				log.Println("quit job")
				ticker.Stop()
				return
			}

		}
	}()
}

func (this *JobExecuter) AddJob(job *Job) {
	if job != nil {
		this.jobs = append(this.jobs, job)
	}
}

func (this *JobExecuter) run() {
	for index, job := range this.jobs {
		log.Println(fmt.Sprintf("start run %d job[%s]", index, job.Name))
		job.Handler()
		log.Println(fmt.Sprintf("finish run %d job[%s]", index, job.Name))

	}
}

func (this *JobExecuter) Stop() {
	if this.quit != nil {
		this.quit <- 1
	}

}
