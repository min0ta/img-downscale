package routinespool

import (
	"fmt"
	"image_thumb/internal/config"
	"image_thumb/internal/db"
	imgproccessor "image_thumb/internal/img_proccessor"
	"time"
)

type Pool struct {
	cfg      *config.Config
	taskChan chan *db.Task
	workers  []*Worker
	proc     *imgproccessor.ImageProcessor
	db       *db.Storage
	launched bool
}

func New(cfg *config.Config, db *db.Storage) *Pool {
	proc := imgproccessor.New(cfg)
	return &Pool{
		cfg:  cfg,
		db:   db,
		proc: proc,
	}
}

func (p *Pool) Start() {
	if p.launched {
		return
	}
	p.launched = true
	for i := 0; i < p.cfg.MaxRoutines; i++ {
		p.workers = append(p.workers, CreateWorker(p.db, p.proc))
	}
	p.taskChan = make(chan *db.Task, 100)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			tasks, err := p.db.GetPending()
			if err != nil {
				fmt.Println("couldnt get pending requests from db wtih error: ", err)
				continue
			}

			for i := 0; i < len(tasks); i++ {
				task := tasks[i]
				p.taskChan <- task
			}
			p.execAll()
		}
	}()
}

func (p *Pool) execAll() {
	i := 0
	for {
		task := p.getTask()
		if task == nil {
			break
		}
		p.workers[i%len(p.workers)].Append(task)
		i++
	}
	for j := 0; j < len(p.workers); j++ {
		go p.workers[j].ExecAll()
	}
}

func (p *Pool) getTask() *db.Task {
	select {
	case task := <-p.taskChan:
		return task
	default:
		return nil
	}
}
