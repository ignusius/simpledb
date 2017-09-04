package simpledb

import (
	"fmt"
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
	//db.Close()

}

func TestPrepare(t *testing.T) {

	err := db.Prepare("INSERT INTO data values ($1,$2,$3,$4,$5)")
	if err != nil {
		t.Fatal("Error! Error Prepare!")
	}
	err = db.StmtExec(1, "test", "test", 2, 3)
	if err != nil {
		t.Fatal("Error! Error Stmt!")
	}
	arr, err := db.Query("SELECT * FROM data")
	if err != nil {
		t.Fatal("Error! Error Query!")
	}
	expected := []interface{}{1, "test", "test", 2, 3}
	if !compare(expected, arr[0]) {
		t.Fatal("Warring! value != expected")

	}
	err = db.Exec("DELETE FROM DATA")
	if err != nil {
		t.Fatal("Error! Clearing database error")
	}

}

func TestTxPrepare(t *testing.T) {

	db.Begin()

	err := db.TxPrepare("INSERT INTO data values ($1,$2,$3,$4,$5)")
	if err != nil {
		t.Fatal("Error! Error Prepare!")
	}
	err = db.StmtExec(1, "test", "test", 2, 3)
	if err != nil {
		t.Fatal("Error! Error Stmt!")
	}
	db.Commit()
	arr, err := db.Query("SELECT * FROM data")
	if err != nil {
		t.Fatal("Error! Error Query!")
	}
	expected := []interface{}{1, "test", "test", 2, 3}
	if !reflect.DeepEqual(expected, arr[0]) {
		t.Fatal("Warring! value != expected")

	}
	//err = db.Exec("DELETE FROM DATA")
	///if err != nil {
	//	t.Fatal("Error! Clearing database error")
	//}

}

func TestClose(t *testing.T) {
	db.Close()
}

func compare(a, b []interface{}) bool {
	fmt.Println(reflect.TypeOf(a[0]))
	fmt.Println(reflect.TypeOf(b[0]))
	fmt.Println(reflect.DeepEqual(a, b))
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if reflect.TypeOf(a[i]) == reflect.Int64 {

			if int64(a[i]) != int64(b[i]) {
				return false

			}

		} else {
			a[i] != b[i]
			return false

		}
		return true
	}
}
