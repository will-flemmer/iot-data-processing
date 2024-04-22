package db

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestGetDBHandle(t *testing.T) {
	dropDB()
	db := GetDBHandle()
	if reflect.TypeOf(db) != reflect.TypeOf((*sql.DB)(nil)) {
		t.Fatal("Setup should return a db handle")
	}
	is_closed_err := db.Ping()
	if is_closed_err != nil {
		t.Fatal(is_closed_err)
	}
}

func TestGetSensorId(t *testing.T) {
	dropDB()
	Setup()
	methods := NewDbMethods()
	id := methods.GetSensorId(FIRST_SENSOR_SERIAL_NUMBER)
	if id != 1 {
		t.Fatalf("Could not fetch id of sensor")
	}
}