package pool

/*
	线程池的概念描述
	ResourceBuffer  chan Resource
	ThreadPool		chan chan Resource

	resource:= <- ResourceBuffer
		thread:= <- ThreadPool
		thread <- resource //获取资源

	DO: //开始工作
		for
			select
				ThreadPool <- aworker
				case resource := <- aworker
					resource.Usage()
*/

var MAXPOOL = 100
var Buff chan Job

type Job struct {
	Task func()
}
type Worker struct {
	WorkerPool chan chan Job
	doJobchan  chan Job
}

func SubmitFunc(task func()) {
	newJob := Job{Task: task}
	Buff <- newJob
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.doJobchan
			select {
			case job := <-w.doJobchan:
				job.Task()
			}
		}
	}()
}
func newWorker(workerpool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerpool,
		doJobchan:  make(chan Job),
	}
}

func initpool() {
	// 初始化 Buff
	Buff = make(chan Job, MAXPOOL)

	pool := make(chan chan Job, MAXPOOL)

	//创建 MAXPOOL 个工人
	for i := 0; i < MAXPOOL; i++ {
		aworker := newWorker(pool)
		aworker.Start()
	}
	//监听buff上 有没有新任务
	go func() {
		for {
			select {
			case job := <-Buff: //获取新任务
				go func(job Job) {
					//获取新工人
					worker := <-pool
					worker <- job
				}(job)
			}
		}
	}()
}

func init() {
	initpool()
}
