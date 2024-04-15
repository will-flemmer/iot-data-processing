package broker

import (
	"encoding/json"
	"fmt"
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
		go handleRequest(conn)
	}
}

type SensorData struct {
	Datetime string  `json:"datetime"`
	TempC    float32 `json:"temp_c"`
}

type SensorDatafile struct {
	SensorId string `json:"sensor_id"`
	Data     []SensorData `json:"data"`
}

// This will fire the various jobs
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	var input SensorDatafile

	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	err = json.Unmarshal(buf[:length], &input)
	if err != nil {
		log.Panic("could not unmarshal data:", err.Error())
	}

	// 1. insert data into db
	// 2. kick off jobs

	returnString := fmt.Sprintf("Data saved, jobs started: %s", buf[:int(length)])
	conn.Write([]byte(returnString))
	conn.Close()
}
