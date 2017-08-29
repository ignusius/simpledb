package simpledb

//SQLstrconv - converting interface{}(SQL string) to string
func SQLstrconv(input interface{}) string {
	return string(input.([]uint8))
}
