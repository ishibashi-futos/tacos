package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Databases int

const (
	MSSQL Databases = iota
	SQLite
	MySQL
)

func (db Databases) String() string {
	switch db {
	case MSSQL:
		return "mssql"
	case SQLite:
		return "sqlite3"
	case MySQL:
		return "mysql"
	default:
		return ""
	}
}

const (
	server = "mssql"
	user = "sa"
	password = "P@s5w0rd"
	port = "1433"
	database = "tacos"
)

func DatabaseFactory(dbType Databases) (db *sql.DB, err error) {
	var connString string
	switch dbType {
	case MSSQL:
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", server, user, password, port, database)
	case SQLite:
		connString = "./thanks.db"
	case MySQL:
		connString = fmt.Sprintf("%s:%s@/%s", server, user, database)
	default:
		return nil, errors.New("")
	}
	db, err = sql.Open(dbType.String(), connString)
	return db, err
}