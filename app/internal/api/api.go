package api

import (
	"encoding/json"
	"fmt"
	"image_thumb/internal/config"
	"image_thumb/internal/db"
	routinespool "image_thumb/internal/routinesPool"
	"io"
	"net/http"
)

type Server struct {
	db   *db.Storage
	cfg  *config.Config
	pool *routinespool.Pool
}

func New(config *config.Config) (*Server, error) {
	db := db.CreatePool(config)
	err := db.Connect()
	if err != nil {
		panic(fmt.Sprintf("cannot connect to db due to %v error", err))
	}
	fmt.Println("succesfully connected to db")
	pool := routinespool.New(config, db)
	pool.Start()
	return &Server{
		db:   db,
		cfg:  config,
		pool: pool,
	}, nil
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", s.addTask)
	mux.HandleFunc("/tasks/{id}", s.getTask)

	return http.ListenAndServe(s.cfg.Port, mux)
}

func (s *Server) getJson(r *http.Request, output any) error {
	text, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(text, output)
	return err
}

func (s *Server) writeJson(w http.ResponseWriter, result any, code int) error {
	w.WriteHeader(code)
	json, err := json.Marshal(result)
	if err != nil {
		return err
	}
	_, err = w.Write(json)
	return err
}

func (s *Server) writeError(w http.ResponseWriter, err string, code int) {
	errMsg := ErrorResponse{
		Error_msg: err,
	}
	er := s.writeJson(w, errMsg, code)
	if er != nil {
		fmt.Println("errrr ", err)
	}
}
