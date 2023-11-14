package task

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)
    
type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID uuid.UUID
	Name string
	State State
	Image string
	Memory int 
	Disk int
	ExposedPorts nat.PortSet
	PortBindings map[string]string
	RestartPolicy string
	StartTime time.Time
	FinishTime time.Time
}

type TaskEvent struct {
	ID uuid.UUID
	State State
	TimeStamp time.Time
	Task Task
}

type TaskConfig struct {
	Name string
	AttachStdInput bool
	AttachStdOutput bool
	AttachStdError bool
	Cmd []string
	Image string
	Memory int64
	Disk int64
	Env []string
	RestartPolicy string
}

type DockerClient struct {
	Client *client.Client
	TaskConfig TaskConfig
	ContainerId string
}
// this is a return value for all commands that would be initiated on the container
type DockerResult struct {
	Error error
	Action string
	ContainerId string
	Result string
}

func (d *DockerClient) Run() DockerResult {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(ctx, d.TaskConfig.Image, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.TaskConfig)
		return DockerResult{Error: err}
	}
	io.Copy(os.Stdout, reader)

	rp := container.RestartPolicy{
		Name: d.TaskConfig.Name,
	}
	r := container.Resources{
		Memory: d.TaskConfig.Memory,
	}
	cc := container.Config{
		Image: d.TaskConfig.Image,
		Env: d.TaskConfig.Env,
	}
	hc := container.HostConfig{
		RestartPolicy: rp,
		Resources: r,
	}
	resp, err:= d.Client.ContainerCreate(ctx, &cc, &hc, nil, nil, d.TaskConfig.Name)
	if err  != nil {
		log.Printf("Error creating container using image %s: %v\n", d.TaskConfig.Name, err)
	}
	err = d.Client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("Error starting container %s: %v\n", d.TaskConfig.Image, err)
		return DockerResult{Error: err}
	}
	d.ContainerId = resp.ID
	out, err := d.Client.ContainerLogs (
		ctx, 
		resp.ID, 
		types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true},
	)
	if err != nil {
		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}

	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out) 
	return DockerResult{
		ContainerId: resp.ID,
		Action: "Start", 
		Result: "Success",
	}
}

func (d *DockerClient) Stop() DockerResult {
	log.Printf("Attempting to stop container %v", d.ContainerId)
	ctx := context.Background()
	err := d.Client.ContainerStop(ctx, d.ContainerId, container.StopOptions{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = d.Client.ContainerRemove(ctx, d.ContainerId, types.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
	
	}
	return DockerResult{Action: "stop", Result: "success", Error: err, ContainerId: id}
}
