package main

import (
	"fmt"
	"iot-data-processing/broker"
	"iot-data-processing/http_server"
)

type TempDataPoint struct {
	temp_c float64
}

type Simulator struct {
	id   string
	data []TempDataPoint
}

func (simulator Simulator) Print() string {
	return fmt.Sprintf("id: %s, last_temp: %v", simulator.id, simulator.data[len(simulator.data)-1].temp_c)
}

func main() {
	go http_server.StartHttpServer()
	go broker.StartBroker()
	var out string
	fmt.Scanln(&out)
}
