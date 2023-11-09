package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/gre8t/cube/task"
)

type Manager struct {
	Pending queue.Queue
	TaskDb map[string][]Task
	EventDb map[string][]TaskEvent
	Workers []string
	WorkersTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}
func (m *Manager) SelectWorker(){
	fmt.Println("Will select an appropriate worker")
}
func (m *Manager) UpdateTasks(){
	fmt.Println("Will update task")
}
func (m *Manager) SendWork() {
	fmt.Println("Will send work to workers")
}