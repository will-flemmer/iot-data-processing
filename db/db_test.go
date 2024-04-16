package db

import (
	"database/sql"
	"fmt"
	"iot-data-processing/types"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	exitVal := m.Run()
	os.Exit(exitVal)
}

func dropDB() {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	dbFile := fmt.Sprintf("%s/test_sql.db", basepath)
	println(dbFile)
	os.Remove(dbFile)
}

func TestSetupReturnsDbHandle(t *testing.T) {
	dropDB()
	db_handle := Setup()
	if reflect.TypeOf(db_handle) != reflect.TypeOf((*sql.DB)(nil)) {
		t.Fatal("Setup should return a db handle")
	}

	is_closed_err := db_handle.Ping()
	if is_closed_err == nil {
		t.Fatal("Database connection is still open, it should be closed")
	}
}

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

func TestSetupCreatesCorrectTables(t *testing.T) {
	dropDB()
	Setup()
	db := GetDBHandle()
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	expectedTableNames := []string{"sensors", "sensor_data", "sqlite_sequence"}
	var tableNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		tableNames = append(tableNames, name)
	}

	slices.Sort(expectedTableNames)
	slices.Sort(tableNames)
	if !reflect.DeepEqual(expectedTableNames, tableNames) {
		t.Fatalf("Unexpected table names. received table names: %v", tableNames)
	}
}


func TestInsertSensorData(t *testing.T) {
	dropDB()
	Setup()
	db := GetDBHandle()
	d := []types.SensorData{
		{ Datetime: "2024-04-14T10:38:30.622Z", TempC: 23.2 },
		{ Datetime: "2024-04-14T10:39:30.622Z", TempC: 26.3 },
	}
	s := types.SensorDatafile{ SensorId: "123", Data: d }
	InsertSensorData(&s, db)

	rows, err := db.Query("SELECT datetime,temp_c,sensor_serial_number FROM sensor_data ORDER BY datetime;")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	// var storedValues []broker.SensorData
	var allRows []types.SensorDataDbRow
	for rows.Next() {
		var r types.SensorDataDbRow
		rows.Scan(&r.Datetime, &r.TempC, &r.SensorId)
		allRows = append(allRows, r)
	}
	fmt.Printf("%v", allRows)

	for i := 0; i < len(allRows); i++ {
		r := allRows[i]
		expect := d[i]
		if r.Datetime != expect.Datetime || r.TempC != expect.TempC || r.SensorId != s.SensorId {
			t.Fatalf("DB rows did not match inputs")
		}
	}
}

