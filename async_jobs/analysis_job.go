package async_jobs

import "iot-data-processing/types"

// aggregate_metrics table
// max_temp
// min_temp
// avg_temp

func AnalysisJob(sensorData *types.SensorDatafile) {
	// calculates aggregate metrics and updates them in db if necessary
	println("running analysis job")
}