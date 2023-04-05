package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	// TODO:: configure a user and don't use root.
	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: "buddhabrot",
	}

	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func update(id int64, filename string) error {
	var err error

	var db *sql.DB
	db, err = connect()
	if err != nil {
		return err
	}

	result, err := db.Exec("UPDATE plots SET pngfile = ? WHERE id = ?",
		filename, id)
	if err != nil {
		return err
	}
	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows != 1 {
		return fmt.Errorf("update affected %d rows", rows)
	}
	return nil
}
