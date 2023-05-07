// Package database provides CRUD operations for the MySQL Buddhabrot database.
package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect connects GORM to MySQL and returns a pointer to a GORM DB and any
// error that occurred.
func Connect() (*gorm.DB, error) {
	const (
		user     = "root"
		database = "buddhabrot"
		passEnv  = "MYSQL_ROOT_PASSWORD"
		server   = "mysql"
	)
	password := os.Getenv(passEnv)
	// Ask the driver to scan DATETIME automatically. See:
	// https://stackoverflow.com/a/45040724/2382333
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password,
		server, database)

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
