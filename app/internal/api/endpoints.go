package api

import (
	"fmt"
	"net/http"
)

func (s *Server) addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var req RequestAddTask
	err := s.getJson(r, &req)
	if err != nil {
		s.writeError(w, ErrorBadRequest, 400)
		return
	}
	task, err := s.db.AddTask(req.Original_url)
	if err != nil {
		fmt.Println("couldnt save task to database with error: ", err)
		s.writeError(w, ErrorBadRequest, 500)
		return
	}

	var response ResponseAddTask
	response.Id = task.Id
	response.Original_url = task.Original_url
	response.Status = task.Current_status

	s.writeJson(w, response, 200)
}

func (s *Server) getTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	task, err := s.db.GetTask(id)
	if err != nil {
		s.writeError(w, ErrorBadRequest, 400)
		return
	}

	var response ResponseGetTask
	response.Status = task.Current_status
	response.Errror_msg = task.Error_msg
	response.Id = task.Id
	response.Original_url = task.Original_url
	response.Result_url = task.Result_url

	s.writeJson(w, response, 200)
}
