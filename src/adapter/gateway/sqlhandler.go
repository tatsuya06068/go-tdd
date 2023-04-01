package gateway

type SqlHandler interface {
	Query(string, ...interface{}) (Row, error)
	Execute(string, ...interface{}) (Result, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
