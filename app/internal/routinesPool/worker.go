package routinespool

import (
	"fmt"
	"image_thumb/internal/db"
	imgproccessor "image_thumb/internal/img_proccessor"
)

type Worker struct {
	proc  *imgproccessor.ImageProcessor
	db    *db.Storage
	queue []*db.Task
}

func CreateWorker(storage *db.Storage, proc *imgproccessor.ImageProcessor) *Worker {
	return &Worker{
		queue: make([]*db.Task, 0),
		db:    storage,
		proc:  proc,
	}
}

func (w *Worker) Append(task *db.Task) {
	w.queue = append(w.queue, task)
}

func (w *Worker) ExecAll() {
	for len(w.queue) > 0 {
		w.exec()
	}
}

func (w *Worker) exec() {
	task := w.pop()

	img, err := w.proc.GetThumbnail(task.Original_url)
	if err != nil {
		fmt.Println("couldnt get thumbnail from original url with error: ", err)
		err := w.db.SetError(task, fmt.Sprintf("Couldnt downloand image due to %v error", err))
		fmt.Printf("couldnt set task status (task.Id = %v) as error with error: %v\n", task.Id, err)
		return
	}
	err = w.proc.SaveThumbnail(img, task.Id)
	if err != nil {
		fmt.Println("couldnt save thumbnail with error: ", err)
		err := w.db.SetError(task, fmt.Sprintf("Couldnt save image due to %v error", err))
		if err != nil {
			fmt.Printf("couldnt set task status (task.Id = %v) as error with error: %v\n", task.Id, err)
		}
	}
	err = w.db.SetDone(task, "http://localhost:8080/storage/"+task.Id+".jpg")
	if err != nil {
		fmt.Printf("couldnt set task status (task.Id = %v) as done with error: %v\n", task.Id, err)
	}
}

func (w *Worker) pop() *db.Task {
	last := len(w.queue) - 1
	task := w.queue[last]
	w.queue = w.queue[:last]
	return task
}
