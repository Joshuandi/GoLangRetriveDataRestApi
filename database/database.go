package database

import "database/sql"

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "admin"
	Dbname   = "HttpRequest"
)

var (
	Db  *sql.DB
	Err error
)
