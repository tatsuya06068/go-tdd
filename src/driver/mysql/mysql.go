package mysql

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tatsuya06068/go-tdd/src/adapter/gateway"
	"golang.org/x/xerrors"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() (gateway.SqlHandler, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, xerrors.Errorf("time failed: %w", err)
	}
	c := mysql.Config{
		DBName:    "db",
		User:      "user",
		Passwd:    "password",
		Addr:      "db",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	conn, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		return nil, xerrors.Errorf("db con failed: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, xerrors.Errorf("ping failed: %w", err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler, nil
}

type SqlResult struct {
	Result sql.Result
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (gateway.Result, error) {
	var res = SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return nil, xerrors.Errorf("query execute failed: %w", err)
	}
	res.Result = result
	return res.Result, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (gateway.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	if err := r.Rows.Scan(dest...); err != nil {
		return err
	}
	return nil
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	if err := r.Rows.Close(); err != nil {
		return err
	}
	return nil
}
