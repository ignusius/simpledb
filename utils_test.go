package simpledb

import "testing"

func TestSQLiteStrconv(t *testing.T) {
	expected := "test"
	var input interface{}
	input = []uint8{116, 101, 115, 116}
	res := sqliteStrconv(input)

	if res != expected {
		t.Fatal("Warring! value != expected")
	}
}
