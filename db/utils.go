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

type DBMethods struct {
	dbHandle *sql.DB
}

func NewDbMethods() *DBMethods {
	db := DBMethods{
		dbHandle: GetDBHandle(),
	}
	return &db
}


func (db DBMethods) InsertSensorData(data *types.SensorDatafile) {
	tx, err := db.dbHandle.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}

	sql := fmt.Sprintf("INSERT INTO %s(datetime, temp_c, sensor_id) VALUES(?, ?, ?)", SENSOR_DATA_TABLE)
	stmt, err := db.dbHandle.Prepare(sql)

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


func (db *DBMethods) GetSensorId(sensor_serial_number string) int32 {
	sql := fmt.Sprintf("SELECT id FROM %s WHERE serial_number='%s' LIMIT 1;", SENSOR_TABLE, sensor_serial_number)
	row := db.dbHandle.QueryRow(sql)

	var id int32
	err := row.Scan(&id)

	if err != nil {
		log.Panicf(err.Error())
	}

	return id
}