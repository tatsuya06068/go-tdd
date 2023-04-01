package repository

import "github.com/tatsuya06068/go-tdd/src/domain/entity"

type TaskRepository interface {
	GetTask() (*entity.Task, error)
}
