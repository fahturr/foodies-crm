package database

import "database/sql"

type Database struct {
	DbSql *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) SetDbSql(dbsql *sql.DB) *Database {
	d.DbSql = dbsql
	return d
}
