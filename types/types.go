package types

type SensorData struct {
	Datetime string  `json:"datetime"`
	TempC    float32 `json:"temp_c"`
}

type SensorDatafile struct {
	SensorId string `json:"sensor_id"`
	Data     []SensorData `json:"data"`
}

type SensorDataDbRow struct {
	TempC float32
	Datetime string
	SensorId string
}
