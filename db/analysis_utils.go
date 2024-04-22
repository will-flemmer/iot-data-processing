package db

import (
	"fmt"
	"log"
)

func (db *DBMethods) GetMaxTemp(sensor_serial_number string) int32 {
	sensor_id := db.GetSensorId(sensor_serial_number)
	stmt := fmt.Sprintf("SELECT MAX(temp_c) FROM %s WHERE sensor_id='%d';", SENSOR_DATA_TABLE, sensor_id)
	row := db.dbHandle.QueryRow(stmt)

	var maxTemp int32
	err := row.Scan(&maxTemp)

	if err != nil {
		log.Fatalf(err.Error())
	}
	return maxTemp
}