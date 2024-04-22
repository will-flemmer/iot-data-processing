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
	AGGREGATE_METRICS_TABLE = "aggregate_metrics"
	FIRST_SENSOR_SERIAL_NUMBER = "1234"
)

func Setup() *sql.DB {
	db := GetDBHandle()
	defer db.Close()

	createSensorTable(db)
	createSensorDataTable(db)
	createAggregateMetricsTabl(db)
	createFirstSensor(db)

	return db
}

func createAggregateMetricsTabl(db_handle *sql.DB) {
	createAggregateMetricsTableStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		max_temp_c DECIMAL(3, 1) NOT NULL,
		min_temp_c DECIMAL(3, 1) NOT NULL,
		avg_temp_c DECIMAL(3, 1) NOT NULL,
		sensor_id INTEGER NOT NULL
	);`, AGGREGATE_METRICS_TABLE)

	_, err := db_handle.Exec(createAggregateMetricsTableStmt)
	if err != nil {
			log.Fatalf("%q: %s\n", err, createAggregateMetricsTableStmt)
	}
}

func createSensorDataTable(db_handle *sql.DB) {
	createSensorDataTableStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		datetime DATETIME NOT NULL,
		temp_c DECIMAL(3, 1) NOT NULL,
		sensor_id INTEGER NOT NULL
	);`, SENSOR_DATA_TABLE)

	_, err := db_handle.Exec(createSensorDataTableStmt)
	if err != nil {
			log.Fatalf("%q: %s\n", err, createSensorDataTableStmt)
	}
}


func createSensorTable(db_handle *sql.DB) {
	createSensorTableStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		serial_number CHAR(8) NOT NULL
);`, SENSOR_TABLE)

	_, err := db_handle.Exec(createSensorTableStmt)
	if err != nil {
			log.Fatalf("%q: %s\n", err, createSensorTableStmt)
	}
}

func createFirstSensor(dbHandle *sql.DB) {
	sql := fmt.Sprintf("INSERT INTO %s(serial_number) VALUES(?)", SENSOR_TABLE)
	stmt, err := dbHandle.Prepare(sql)

	if err != nil {
			log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(FIRST_SENSOR_SERIAL_NUMBER)
	if err != nil {
		log.Panicf(err.Error())
	}
}
