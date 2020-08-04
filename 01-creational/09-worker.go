package creational

import "fmt"

//JobChan 任务队列
type JobChan chan Job

var (
	jobNum   = 100
	jobQueue = make(JobChan, jobNum) //任务队列
)

// Job 对象
type Job struct {
	Content string
}

// createJob 生成任务队列
func createJob() {
	for i := 0; i < jobNum; i++ {
		job := Job{Content: fmt.Sprintf("第%d任务", i)}
		jobQueue <- job
	}
}

// Worker 模式
type Worker struct {
	WorkerPool chan JobChan // 工人池(所有worker共享)
	WorkChan   JobChan      // 工人
	WorkName   string       // 工人名字
	QuitChan   chan bool    // 退出信号
}

// NewWorker 创建worker
func NewWorker(workerPool chan JobChan, i int) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		WorkChan:   make(JobChan),
		WorkName:   fmt.Sprintf("工人%d完成", i),
		QuitChan:   make(chan bool),
	}
}

// Start 启动
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.WorkChan
			select {
			case job := <-w.WorkChan:
				fmt.Println(w.WorkName + job.Content)
			case <-w.QuitChan:
				return
			}
		}
	}()
}

// Stop 退出
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Dispatcher 调度器
type Dispatcher struct {
	WorkerPool chan JobChan
	MaxWorks   int
}

// NewDispatcher 创建调度器
func NewDispatcher(maxWorks int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan JobChan, maxWorks),
		MaxWorks:   maxWorks,
	}
}

// Run 运行
func (d *Dispatcher) Run() {
	// 创建工人
	for i := 0; i < d.MaxWorks; i++ {
		worker := NewWorker(d.WorkerPool, i+1)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue: //获取任务
			go func(job Job) {
				workerChanel := <-d.WorkerPool //工人池获取一工人
				workerChanel <- job            //任务安排给工人
			}(job)
		}
	}
}
