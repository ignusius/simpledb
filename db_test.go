package simpledb

import (
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db DB

func TestNewDatabase(t *testing.T) {
	err := db.NewDatabase("sqlite3", "db_test.db")
	if err != nil {
		t.Fatal("Error! file db_test.db not found!")
	}
}

func TestExec(t *testing.T) {
	err := db.Exec("INSERT INTO data values ($1,$2,$3,$4,$5)", 1, "test", "test", 2, 3)
	if err != nil {
		t.Fatal("Error! Error INSERT!")
	}

}

func TestQuery(t *testing.T) {

	arr, err := db.Query("SELECT * FROM data")
	if err != nil {
		t.Fatal("Error! Error Query!")
	}
	expected := []interface{}{1, "test", "test", 2, 3}
	if reflect.DeepEqual(expected, arr[0]) {
		t.Fatal("Warring! value != expected")
		
	}
	err = db.Exec("DELETE FROM DATA")
	if err != nil {
		t.Fatal("Error! Clearing database error")
	}
	db.Close()

}
