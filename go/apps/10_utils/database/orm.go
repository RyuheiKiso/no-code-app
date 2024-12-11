package orm

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ORM struct {
	DB *sql.DB
}

func NewORM(dataSourceName string) (*ORM, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &ORM{DB: db}, nil
}

func (o *ORM) Create(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := o.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

func (o *ORM) Read(query string, args ...interface{}) (*sql.Rows, error) {
	return o.DB.Query(query, args...)
}

func (o *ORM) Update(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := o.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

func (o *ORM) Delete(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := o.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

func (o *ORM) BeginTransaction() (*sql.Tx, error) {
	return o.DB.Begin()
}

func (o *ORM) RollbackTransaction(tx *sql.Tx) error {
	return tx.Rollback()
}

func (o *ORM) CommitTransaction(tx *sql.Tx) error {
	return tx.Commit()
}

func (o *ORM) Close() error {
	return o.DB.Close()
}
