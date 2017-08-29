package simpledb

func SQLstrconv(input interface{}) string {
	return string(input.([]uint8))
}
