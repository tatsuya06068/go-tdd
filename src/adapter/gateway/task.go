package gateway

import (
	"database/sql"

	"github.com/tatsuya06068/go-tdd/src/domain/model"
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

func (t *TaskRepository) GetTask() (*model.Task, error) {
	conn := t.conn
	row := conn.QueryRow("SELECT * FROM 'task' ")
	task := model.Task{}
	err := row.Scan(&task.ID, &task.TaskName, &task.Status)
	if err != nil {
		return nil, err
	}

	return &task, nil

}
