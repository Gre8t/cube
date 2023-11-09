package main

import (
	"github.com/gre8t/cube/node"
	"github.com/gre8t/cube/task"
	"fmt"
	"time"

	"github.com/gre8t/cube/manager"
	"github.com/gre8t/cube/worker"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	t := task.Task{
		ID:   uuid.New(),
		Name: "Task-1",
		State: task.Pending,
		Image: "Image-1",
		Memory: 1024,
		Disk: 1,
	}
}
