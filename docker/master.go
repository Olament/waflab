// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package docker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	TaskStatusQueue   = 1 // inside work channel, waiting to be execute
	TaskStatusRunning = 2
	TaskStatusFinish  = 3 // finished, ready to be removed from taskStatus
)

var (
	ErrInitDocker    = errors.New("docker/master: failed to initialize Docker client")
	ErrInvalidMaster = errors.New("docker/worker: Master is not properly initialized")

	ctx = context.Background()
	cli *client.Client
)

type taskStatus struct {
	status    int
	numRetry  int
	startTime time.Time
}

type Master struct {
	success  bool
	mux      sync.Mutex
	taskChan chan *Task
	status   map[string]*taskStatus // taskID -> taskStatus
}

// Master worker communication
func (m *Master) getTask() *Task {
	task := <-m.taskChan

	m.mux.Lock()
	defer m.mux.Unlock()

	m.status[task.ID].status = TaskStatusRunning

	return task
}

func (m *Master) reportTask(task *Task, isDone bool, res []Response, err error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if isDone {
		m.status[task.ID].status = TaskStatusFinish
		task.Res <- res
	} else {
		if m.status[task.ID].numRetry < MaxRetryTime {
			m.status[task.ID].status = TaskStatusQueue
			m.status[task.ID].numRetry += 1
			m.taskChan <- task
		} else {
			m.status[task.ID].status = TaskStatusFinish
			task.Err <- err
		}
	}
}

// EXPOSED
func (m *Master) InsertTask(hostname string, yamlFile string) ([]Response, error) {
	if !m.success {
		return nil, ErrInvalidMaster
	}

	resChan := make(chan []Response)
	errChan := make(chan error)
	id := strconv.FormatInt(time.Now().Unix(), 10)
	task := Task{
		ID:       id,
		Hostname: hostname,
		YAMLFile: yamlFile,
		Res:      resChan,
		Err:      errChan,
	}
	m.taskChan <- &task
	m.mux.Lock()
	m.status[id] = &taskStatus{
		status:    TaskStatusQueue,
		numRetry:  0,
		startTime: time.Now(),
	}
	m.mux.Unlock()

	select {
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return nil, err
	}
}

func MakeMaster(numContainer int) *Master {
	m := Master{}

	m.status = make(map[string]*taskStatus)
	m.taskChan = make(chan *Task, 100)

	cli, _ = client.NewClientWithOpts(client.FromEnv)

	if !GetMasterHealth() {
		fmt.Printf("\033[31m%v\nDo you have Docker (https://www.docker.com/) installed?\033[0m\n", ErrInitDocker)
		return &m
	}

	reader, err := cli.ImagePull(ctx, "docker.io/waflab/wafbench", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	startingPort := 17000
	for i := 0; i < numContainer; i++ {
		worker, err := MakeWorker(&m, cli, ctx, strconv.Itoa(startingPort+i))
		if err != nil {
			panic(err)
		}
		go worker.Run()
	}

	return &m
}

func GetMasterHealth() bool {
	if cli == nil { // client is not initialized
		return false
	}
	_, err := cli.Ping(ctx)
	return err == nil
}
