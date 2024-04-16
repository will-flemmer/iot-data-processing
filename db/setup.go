package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SENSOR_TABLE = "sensors"
	SENSOR_DATA_TABLE = "sensor_data"
)

func Setup() *sql.DB {
	db := GetDBHandle()
	defer db.Close()

	createSensorTableStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			serial_number CHAR(8) NOT NULL
	);`, SENSOR_TABLE)

	_, err := db.Exec(createSensorTableStmt)
	if err != nil {
			log.Fatalf("%q: %s\n", err, createSensorTableStmt)
	}

	createSensorDataTableStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		datetime DATETIME NOT NULL,
		temp_c DECIMAL(3, 1) NOT NULL,
		sensor_serial_number CHAR(8) NOT NULL
	);`, SENSOR_DATA_TABLE)

	_, err = db.Exec(createSensorDataTableStmt)
	if err != nil {
			log.Fatalf("%q: %s\n", err, createSensorTableStmt)
	}

	return db
}
