// Package database provides CRUD operations for the MySQL Buddhabrot database.
package database

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

// Insert inserts a plot record with the given JSON Buddhabrot parameters, and
// returns the ID of the newly created plot.
func Insert(json string) (int64, error) {
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

// GetFilename retrieves the filename associated with the given plot ID.
func GetFilename(id int64) (string, error) {
	var db *sql.DB
	db, err := connect()
	if err != nil {
		return "", err
	}

	var filename string
	row := db.QueryRow("SELECT pngfile FROM plots WHERE id = ?", id)
	if row.Scan(&filename); err != nil {
		return "", err
	}

	return filename, nil
}

// Update updates the filename associated with the given plot ID.
func Update(id int64, filename string) error {
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
