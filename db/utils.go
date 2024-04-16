package db

import (
	"database/sql"
	"fmt"
	"log"

	"iot-data-processing/broker"
	"os"

	_ "github.com/mattn/go-sqlite3"
)


func getDBName() string {
	env := os.Getenv("ENV")
	if env == "test" {
		return "db/test_sql.db"
	}
	return "db/dev_sql.db"
}

// It is the responsibility of the calling function to
// close the db handle by doing:
// db.Close()
func GetDBHandle() *sql.DB {
	db, err := sql.Open("sqlite3", getDBName())
	if err != nil {
			log.Fatal(err)
	}
	return db
}


func InsertSensorData(data *broker.SensorDatafile, db *sql.DB, sensor_serial_number string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}

	sql := fmt.Sprintf("INSERT INTO %s(datetime, temp_c, sensor_serial_number) VALUES(?, ?, ?)", SENSOR_TABLE)
	stmt, err := db.Prepare(sql)

	if err != nil {
			log.Fatal(err)
	}
	defer stmt.Close()

	for _, dataPoint := range data.Data {
		// start a transaction?
		_, err = stmt.Exec(dataPoint.Datetime, dataPoint.TempC, data.SensorId)
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			break
		}
	}
		err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
