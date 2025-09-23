package db

import "time"

type Task struct {
	Id             string
	Original_url   string
	Current_status string
	Result_url     string
	Error_msg      string
	Created_at     time.Time
	Updated_at     time.Time
}
