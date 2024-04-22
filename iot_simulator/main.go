package main

import (
	"encoding/json"
	"iot-data-processing/broker"
	"iot-data-processing/types"
	"log"
	"net"
)

func main() {
	// send json to 
	println("hello")

	datafile := types.SensorDatafile{
		SensorId: "123",
		Data: []types.SensorData {
			{ Datetime: "2024-04-14T10:38:30.622Z", TempC: 20.0 },
			{ Datetime: "2024-04-14T10:39:30.622Z", TempC: 20.5 },
			{ Datetime: "2024-04-14T10:40:30.622Z", TempC: 20.5 },
		},
	}
	json_string, err := json.Marshal(datafile)

	if err != nil {
		log.Panic(err)
	}
	println(string(json_string))

	conn, err := net.Dial(broker.CONN_TYPE, broker.CONN_HOST+":"+broker.CONN_PORT)

	if err != nil {
		log.Panic(err)
	}
	conn.Write([]byte(json_string))
}