package broker

import (
	"encoding/json"
	"fmt"
	"iot-data-processing/async_jobs"
	"iot-data-processing/db"
	"iot-data-processing/types"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func StartBroker() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	db_handle := db.NewDbMethods()
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Broker listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn, db_handle)
	}
}


func handleRequest(conn net.Conn, db_handle *db.DBMethods) {
	buf := make([]byte, 1024)
	var sensor_data_file types.SensorDatafile

	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	err = json.Unmarshal(buf[:length], &sensor_data_file)
	if err != nil {
		log.Panic("could not unmarshal data:", err.Error())
	}

	// db_handle.InsertSensorData(&sensor_data_file)

	go async_jobs.AlertingJob(&sensor_data_file, db_handle)
	go async_jobs.AnalysisJob(&sensor_data_file, db_handle)

	returnString := fmt.Sprintf("Data saved, jobs started: %s", buf[:int(length)])
	conn.Write([]byte(returnString))
	conn.Close()
}
