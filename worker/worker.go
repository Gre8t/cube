//this package is responsible for our worker node
package worker
import (
	"fmt"
	"github.com/google/uuid"
	"github.com/golang-collections/collections/queue"

	"cube/task"
)

type Worker struct{
	Name string
	Queue queue.Queue
	Db map[uuid.UUID]task.Task
	TaskCount int
}
func (w *Worker) CollectStats(){
	fmt.Println("I will collect stats")
}   
func (w *Worker) RunTask(){
	fmt.Println("This will start and stop task")
}
func (w *Worker) StartTask(){
	fmt.Println("This will start task")
}
func (w *Worker) StopTask(){
	fmt.Println("This will stop task")
}