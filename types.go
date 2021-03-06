package jet

import (
	"database/sql"
)

type Migration struct {
	Up   string
	Down string
	Id   int64
}

type Db interface {
	Queryable
	// Begin starts a transaction
	Begin() (Tx, error)

	// SetColumnConverter sets the converter instance to use
	// when converting from db column names to struct fields.
	SetColumnConverter(conv ColumnConverter)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//
	// If n <= 0, no idle connections are retained.
	SetMaxIdleConns(n int)

	// ExpandMapAndSliceMarker expands a marker to several if argument is a slice or map.
	ExpandMapAndSliceMarker(f bool)
}

type Tx interface {
	Queryable
	// Commit commits the transaction
	Commit() error
	// Rollback rolls back the transaction
	Rollback() error
}

type Queryable interface {
	// Query prepares the query for execution
	Query(query string, args ...interface{}) Queryable
	// Run runs the query without returning results
	Run() error
	// Rows runs the query writing the rows to the specified map or struct array. If maxRows is specified, only writes up to maxRows rows.
	Rows(v interface{}, maxRows ...int64) error
	// Logger returns the current logger
	Logger() *Logger
	// SetLogger sets a logger
	SetLogger(l *Logger)
}

type queryObject interface {
	Prepare(query string) (*sql.Stmt, error)
}
