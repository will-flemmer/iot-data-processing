package async_jobs

import (
	"iot-data-processing/db"
	"iot-data-processing/types"
)

// aggregate_metrics table
// max_temp
// min_temp
// avg_temp
// sensor_id

func AnalysisJob(sensorDataFile *types.SensorDatafile, db *db.DBMethods) {
	// calculates aggregate metrics and updates them in db if necessary
	println("running analysis job")
	updateMaxTemp(sensorDataFile)
}

func updateMaxTemp(sensorDataFile *types.SensorDatafile) {
	dbM := db.NewDbMethods()
	maxTemp := dbM.GetMaxTemp(db.FIRST_SENSOR_SERIAL_NUMBER)
	println(maxTemp)
	// fetch max temp for sensor

	// update if necessary
}