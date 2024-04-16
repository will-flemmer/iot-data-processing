package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"iot-data-processing/types"

	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func getDBName() string {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	env := os.Getenv("ENV")
	if env == "test" {
		return fmt.Sprintf("%s/test_sql.db", basepath)
	}
	return fmt.Sprintf("%s/dev_sql.db", basepath)
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


func InsertSensorData(data *types.SensorDatafile, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}

	sql := fmt.Sprintf("INSERT INTO %s(datetime, temp_c, sensor_serial_number) VALUES(?, ?, ?)", SENSOR_DATA_TABLE)
	stmt, err := db.Prepare(sql)

	if err != nil {
			log.Fatal(err)
	}
	defer stmt.Close()

	for _, dataPoint := range data.Data {
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
