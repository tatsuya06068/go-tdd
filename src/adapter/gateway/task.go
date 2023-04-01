package gateway

import (
	"database/sql"

	"github.com/tatsuya06068/go-tdd/src/domain/entity"
	"github.com/tatsuya06068/go-tdd/src/domain/repository"
)

type TaskRepository struct {
	conn *sql.DB
}

func NewTaskRepository(conn *sql.DB) repository.TaskRepository {
	return &TaskRepository{
		conn: conn,
	}
}

func (t *TaskRepository) GetTask() (*entity.Task, error) {
	conn := t.conn
	row := conn.QueryRow("SELECT * FROM 'task' ")
	task := entity.Task{}
	err := row.Scan(&task.TastId, &task.TaskName)
	if err != nil {
		return nil, err
	}

	return &task, nil

}
