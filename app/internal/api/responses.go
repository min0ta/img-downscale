package api

// response to POST /api/tasks
type ResponseAddTask struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	Original_url string `json:"original_url"`
}

// requst to POST /api/tasks
type RequestAddTask struct {
	Original_url string `json:"original_url"`
}

// response to GET /api/tasks
type ResponseGetTask struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	Original_url string `json:"original_url"`
	Result_url   string `json:"result_url"`
	Errror_msg   string `json:"error_msg"`
}

// response with error
type ErrorResponse struct {
	Error_msg string `json:"error_msg"`
}

const (
	ErrorBadRequest = "bad request"
)
