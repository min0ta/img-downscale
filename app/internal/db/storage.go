package db

import (
	"context"
	"image_thumb/internal/config"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	cfg  *config.Config
	pool *pgxpool.Pool
}

func CreatePool(cfg *config.Config) *Storage {
	db := &Storage{
		cfg: cfg,
	}
	return db
}

func (db *Storage) ExecuteMigrations() error {
	migrationsFile, err := os.ReadFile(db.cfg.MigrationsPath)
	if err != nil {
		return err
	}
	sql := string(migrationsFile)
	_, err = db.pool.Exec(context.Background(), sql)
	return err
}

func (db *Storage) Connect() error {
	pool, err := pgxpool.New(context.Background(), db.cfg.DBConnectionURL)
	if err != nil {
		return err
	}
	db.pool = pool
	return pool.Ping(context.Background())
}

func (db *Storage) SetError(t *Task, err string) error {
	var id string
	return db.pool.QueryRow(context.Background(),
		"UPDATE tasks SET error_msg = $2, updated_at = NOW(), current_status = 'error' WHERE id = $1 RETURNING id",
		t.Id, err,
	).Scan(&id)
}

func (db *Storage) GetPending() ([]*Task, error) {
	var tasks []*Task
	rows, err := db.pool.Query(context.Background(), "SELECT id, original_url, current_status, result_url, error_msg, created_at, updated_at FROM tasks WHERE current_status = 'pending'")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.Id, &task.Original_url, &task.Current_status, &task.Result_url, &task.Error_msg, &task.Created_at, &task.Updated_at)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func (db *Storage) AddTask(url string) (*Task, error) {
	task := &Task{}
	err := db.pool.QueryRow(context.Background(), "INSERT INTO tasks (original_url, current_status) VALUES ($1, 'pending') RETURNING id", url).Scan(&task.Id)
	if err != nil {
		return nil, err
	}
	task.Current_status = "pending"
	task.Original_url = url
	return task, nil
}

func (db *Storage) GetTask(id string) (*Task, error) {
	task := &Task{}
	err := db.pool.QueryRow(context.Background(), "SELECT id, current_status, original_url, original_url, result_url, error_msg FROM tasks WHERE id = $1", id).Scan(&task.Id, &task.Current_status, &task.Original_url, &task.Result_url, &task.Error_msg)
	return task, err
}

func (db *Storage) SetDone(t *Task, result_url string) error {
	var id string
	return db.pool.QueryRow(context.Background(),
		"UPDATE tasks SET result_url = $2, updated_at = NOW(), current_status = 'done' WHERE id = $1 RETURNING id",
		t.Id, result_url,
	).Scan(&id)
}
