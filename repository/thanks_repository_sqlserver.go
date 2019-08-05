package repository

import (
	"database/sql"
)

type SQLServerQueryImpl struct {
	conn *sql.DB
}

func NewSQLServerQueryImpl(db *sql.DB) *SQLServerQueryImpl {
	return &SQLServerQueryImpl{db}
}

func (db *SQLServerQueryImpl) ExecuteSave(from_at string, to_at string, message string) (*sql.Result, error) {
	sql := "INSERT INTO THANKS VALUES(NEWID(), GETDATE(), ?, ?, ?)"
	result, err := db.conn.Exec(sql, from_at, to_at, message);
	if err != nil {
		return nil, err
	}
	return &result, nil
}