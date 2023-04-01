package interator

import (
	"fmt"

	"github.com/tatsuya06068/go-tdd/src/domain/entity"
	"github.com/tatsuya06068/go-tdd/src/domain/repository"
	"github.com/tatsuya06068/go-tdd/src/usecase/port"
)

type TaskInteractor struct {
	TaskOutPutPort port.TaskOutPutPort
	TaskRepository repository.TaskRepository
}

func NewTaskInteractor(taskOutput port.TaskOutPutPort, taskRepository repository.TaskRepository) *TaskInteractor {
	return &TaskInteractor{
		TaskOutPutPort: taskOutput,
		TaskRepository: taskRepository,
	}
}

func (i *TaskInteractor) GetTaskData() (*entity.Task, error) {
	res, err := i.TaskRepository.GetTask()
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
