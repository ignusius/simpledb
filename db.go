//Package simpledb for simple use databases.
package simpledb

import (
	"fmt"
	"database/sql"
	"errors"
)

//DB is structure of database
type DB struct {
	Db   *sql.DB
	err  error
	rows *sql.Rows
}

//NewDatabase - create new database
func (d *DB) NewDatabase(database, basename string) error {
	d.Db, d.err = sql.Open(database, basename)
	if d.err != nil {
		fmt.Println(d.err)
		return errors.New("DataBase not found")
	}
	d.err = d.Db.Ping()
	if d.err != nil {
		fmt.Println(d.err)
		return errors.New("Connection failed")
	}
	return nil
}

//Query - for query from database
func (d *DB) Query(query string, args ...interface{}) ([][]interface{}, error) {
	var emptyReturn [][]interface{}
	d.rows, d.err = d.Db.Query(query, args...)
	if d.err != nil {
		fmt.Println(d.err)
		return emptyReturn, errors.New("Query -> query error")
	}
	defer d.rows.Close()
	var varReturn [][]interface{}

	col, _ := d.rows.Columns()
	varArr := make([]interface{}, len(col))
	varArrPtrs := make([]interface{}, len(col))
	counter := 0

	for d.rows.Next() {
		var tmp []interface{}
		for i := range col {
			varArrPtrs[i] = &varArr[i]
		}
		d.err = d.rows.Scan(varArrPtrs...)
		if d.err != nil {
			fmt.Println(d.err)
			return emptyReturn, errors.New("Query -> row.Scan error")
		}
		varReturn = append(varReturn, tmp)
		for i := range varArr {
			switch t := varArr[i].(type) {

			//In SQLite and MySQL text is []uint8 type
			case []uint8:
			
				varReturn[counter] = append(varReturn[counter], SQLiteStrconv(varArr[i]))

			default:
				_ = t
				varReturn[counter] = append(varReturn[counter], varArr[i])
			}

		}
		counter++
	}

	return varReturn, nil

}

//Exec - for exec to database
func (d *DB) Exec(exec string, args ...interface{}) error {
	_, d.err = d.Db.Exec(exec, args...)
	if d.err != nil {
		fmt.Println(d.err)
		return d.err
	}
	return nil
}

//Close - close  database
func (d *DB) Close() {
	d.Db.Close()
}

