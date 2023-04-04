package main

import (
	"database/sql"
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

func insert(json string) (int64, error) {
	var err error

	var db *sql.DB
	db, err = connect()
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO plots (plot) VALUES (?)", json)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
