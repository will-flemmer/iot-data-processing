package broker

import (
	"database/sql"
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
	db_handle := db.GetDBHandle()
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn, db_handle)
	}
}


func handleRequest(conn net.Conn, db_handle *sql.DB) {
	buf := make([]byte, 1024)
	var input types.SensorDatafile

	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	err = json.Unmarshal(buf[:length], &input)
	if err != nil {
		log.Panic("could not unmarshal data:", err.Error())
	}

	db.InsertSensorData(&input, db_handle)

	go async_jobs.AnalysisJob()

	returnString := fmt.Sprintf("Data saved, jobs started: %s", buf[:int(length)])
	conn.Write([]byte(returnString))
	conn.Close()
}
