package repository

import (
	"database/sql"
	"strings"
)

type Period int
const (
	Unknown Period = iota
	Week
	Month
	Year
)

func (p Period) Names() []string {
	return []string {
		"unknown",
		"Weekly",
		"Monthly",
		"Annual",
	}
}

func (p Period) String() string {
	return p.Names()[p]
}

func NewPeriod(strP string) Period {
	P := strings.ToUpper(strP)
	var p Period
	for i, name := range p.Names() {
		if P == strings.ToUpper(name) {
			return Period(i)
		}
	}
	return Unknown
}

type ThanksRepository struct {
	conn *sql.DB
	query IThanksQuery
}

type IThanksQuery interface {
	ExecuteSave(from_at string, to_at string, message string) (*sql.Result, error)
	// ExecuteSummarize(p Period) (*sql.Rows, error)
}

func NewThanksRepository() (*ThanksRepository, error) {
	var dbType Databases = MSSQL
	db, err := DatabaseFactory(dbType)
	impl := NewSQLServerQueryImpl(db)

	return &ThanksRepository{db, impl}, err
}

func (db *ThanksRepository) Close() {
	db.conn.Close()
}

func (db *ThanksRepository) Save(from_at string, to_at string, message string) error {
	if _, err := db.query.ExecuteSave(from_at, to_at, message); err != nil {
		return err
	}
	return nil
}

func (db *ThanksRepository) Summarize(p Period) (map[string]int, error) {
	result := make(map[string]int)
	sql := "SELECT COUNT(1) AS CNT, to_at FROM THANKS "
	switch p {
	case Week:
		sql += " WHERE CONVERT(varchar, post_date, 112) > DATEADD(day, -7, GETDATE()) AND CONVERT(varchar, post_date, 112) < DATEADD(day, -1, GETDATE())"
	case Month:
		sql += " WHERE LEFT(CONVERT(varchar, post_date, 112), 6) = LEFT(CONVERT(varchar, DATEADD(MONTH, -1, GETDATE()), 112), 6)"
	case Year:
		sql += " WHERE LEFT(CONVERT(varchar, post_date, 112), 4) = LEFT(CONVERT(varchar, DATEADD(MONTH, -1, GETDATE()), 112), 4)"
	}
	sql += " GROUP BY TO_AT"
	rows, err := db.conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var to_at string
		var cnt int
		if err = rows.Scan(&cnt, &to_at); err != nil {
			return nil, err
		}
		result[to_at] = cnt
	}
	return result, nil
}