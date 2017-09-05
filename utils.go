package simpledb

import (
	"reflect"
)

//SQLiteStrconv - converting interface{}(SQL string) to string
func sqliteStrconv(input interface{}) string {
	return string(input.([]uint8))
}

//compare interface{} for testing
func compare(a, b []interface{}) bool {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if reflect.TypeOf(b[i]).String() == "int64" {
			if a[i] != int(b[i].(int64)) {
				return false
			}
		} else {
			if a[i] != b[i] {
				return false
			}
		}

	}
	return true
}
