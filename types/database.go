package types

import (
	"database/sql"
)

// NewDatabaseDriver is a wrapper to create a new database driver
func NewDatabaseDriver(db *sql.DB) Database {
	return Database{conn: db}
}

// Database is a wrapper around database/sql for unit tests
type Database struct {
	conn *sql.DB
}

func (receiver Database) Query() {
	panic("implement me")
}

// Close closes the database connection
func (receiver Database) Close() error {
	return receiver.conn.Close()
}

// Exec executes a query
func (receiver Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return receiver.conn.Exec(query, args...)
}

// Begin begins a transaction query
func (receiver Database) Begin() (TransactionDriver, error) {
	trx, err := receiver.conn.Begin()

	if err != nil {
		return nil, err
	}

	return Transaction{
		trx: trx,
	}, nil
}

// Transaction is a wrapper around sql.tx for mocking and unit tests
type Transaction struct {
	trx *sql.Tx
}

// Prepare is a method that prepares for a transactional query
func (t Transaction) Prepare(query string) (StatementDriver, error) {
	stmt, err := t.trx.Prepare(query)
	if err != nil {
		return nil, err
	}
	return Statement{
		stmt: stmt,
	}, nil
}

// Commit commits the query
func (t Transaction) Commit() error {
	return t.trx.Commit()
}

// Statement is a wrapper for mocking and unit tests
type Statement struct {
	stmt StatementDriver
}

// Exec executes the statement with params
func (s Statement) Exec(args ...interface{}) (sql.Result, error) {
	return s.stmt.Exec(args...)
}

// Close closes the statement query
func (s Statement) Close() error {
	return s.stmt.Close()
}

// StatementDriver is an interface for a transaction statement
type StatementDriver interface {
	Close() error
	Exec(args ...interface{}) (sql.Result, error)
}

// TransactionDriver is an interface for database transactions
type TransactionDriver interface {
	Commit() error
	Prepare(query string) (StatementDriver, error)
}

// DatabaseDriver is an interface so mocking unit tests is possible
type DatabaseDriver interface {
	Close() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Begin() (TransactionDriver, error)
	Query()
}
