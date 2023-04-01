package port

import "github.com/tatsuya06068/go-tdd/src/domain/entity"

/*
* Input Port : Interactorで実装、Controllerで使用される
 */
type TaskInputPort interface {
	GetTaskData() (*entity.Task, error)
}

/*
* OutPut Port : Presenterで実装、Interactorで使用される
 */
type TaskOutPutPort interface {
	GetTaskData([]entity.Task, error)
}
