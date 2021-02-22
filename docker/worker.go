// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package docker

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	scheduleInterval = time.Millisecond * 50
)

type Worker struct {
	mux          sync.Mutex
	master       *Master
	dockerClient *client.Client
	httpClient   *http.Client
	context      context.Context
	port         string
	numJob       uint64 // number of ongoing jobs for this worker
	jobCapacity  uint64 // max number of jobs
}

func (w *Worker) Run() {
	for {
		if w.numJob < w.jobCapacity {
			fmt.Printf("%s start find job with %v ongoing jobs\n", w.port, w.numJob)
			task := w.master.getTask()
			fmt.Printf("%s get job %s\n", w.port, task.ID)
			go w.doTask(task)
		}
		time.Sleep(scheduleInterval)
	}
}

func (w *Worker) doTask(task *Task) {
	fmt.Printf("Start task %s on %s\n", task.ID, w.port)
	atomic.AddUint64(&w.numJob, 1)
	defer atomic.AddUint64(&w.numJob, ^uint64(0)) // decremnt by one

	resp, err := sendRequest(w.httpClient,
		fmt.Sprintf("http://127.0.0.1:%s", w.port),
		task.Hostname,
		task.YAMLFile)
	if err != nil {
		w.master.reportTask(task, false, nil, err)
		return
	}
	defer resp.Body.Close()

	var result []Response
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	fmt.Println(result)
	if err != nil {
		w.master.reportTask(task, false, nil, err)
	}

	w.master.reportTask(task, true, result, nil)
}

// https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data
func sendRequest(client *http.Client, url string, target string, file string) (res *http.Response, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	values := map[string]io.Reader{
		"file":     strings.NewReader(file),
		"hostname": strings.NewReader(target),
	}

	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if key == "file" {
			if fw, err = w.CreateFormFile(key, "temp.yaml"); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}

	w.Close()

	req, err := http.NewRequest("GET", url, &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return client.Do(req)
}

func MakeWorker(master *Master, cli *client.Client, ctx context.Context, port string) (*Worker, error) {
	w := Worker{}
	w.mux.Lock()
	defer w.mux.Unlock()

	fmt.Println(port)

	containerName := fmt.Sprintf("wafbench-%s", port)
	containerID, err := getContainerByName(ctx, cli, containerName)
	if err != nil {
		return nil, err
	}

	// create a new container if there is not one already
	if containerID == "" {
		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image:      "olament/wafbench",
			WorkingDir: "/WAFBench/ftw_compatible_tool",
			ExposedPorts: nat.PortSet{
				"5000": struct{}{},
			},
			Cmd: []string{
				"gunicorn",
				"--bind", "0.0.0.0:5000",
				"web_interface:app",
			},
			Tty: true,
		},
			&container.HostConfig{
				PortBindings: nat.PortMap{
					"5000": []nat.PortBinding{
						{
							HostIP:   "0.0.0.0",
							HostPort: port,
						},
					},
				},
			}, nil, containerName)
		if err != nil {
			return nil, err
		}
		containerID = resp.ID
	}

	// restart the container
	zeroDuration := time.Since(time.Now())
	cli.ContainerRestart(ctx, containerID, &zeroDuration)
	if err != nil {
		return nil, err
	}

	// wait for gunicorn to start
	out, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
		Since:      strconv.Itoa(int(time.Now().UTC().Unix())),
	})
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(out)
	for scanner.Scan() { // blocking until gunicorn start
		if strings.Contains(scanner.Text(), "Booting worker") {
			break
		}
	}

	// http client
	httpClient := http.Client{}

	w.master = master
	w.dockerClient = cli
	w.httpClient = &httpClient
	w.context = ctx
	w.port = port
	w.numJob = 0
	w.jobCapacity = MaxJobPerWorker

	return &w, nil
}

func getContainerByName(ctx context.Context, client *client.Client, name string) (string, error) {
	containerList, err := client.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return "", err
	}
	for _, instance := range containerList {
		for _, partialName := range instance.Names {
			if strings.Contains(partialName, name) {
				return instance.ID, nil
			}
		}
	}
	return "", nil
}
