package async_jobs

import (
	"iot-data-processing/db"
	"iot-data-processing/types"
)

// alerts table:
// sensor_id
// alert_id

// alert_types table:
// message
// code
// --- code --- | --- message ---
// temp_high		| temperature is above 25 degrees c
// temp_low			| temperature is below 20 degrees c

func AlertingJob(sensorData *types.SensorDatafile, db *db.DBMethods) {
	println("running alerting job")
	// check if any alert conditions are met
	// if so, create alert in db
}