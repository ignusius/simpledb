package simpledb

//SQLiteStrconv - converting interface{}(SQL string) to string
func SQLiteStrconv(input interface{}) string {
	return string(input.([]uint8))
}
