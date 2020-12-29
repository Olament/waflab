package docker

const MaxJobPerWorker = 4
const MaxRetryTime = 0

type Task struct {
	ID string
	Hostname string
	YAMLPath string
	Res chan []Response
	Err chan error
}

type Response struct {
	Title string `json:"title"`
	Status []int `json:"status"`
}