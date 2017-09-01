//Package simpledb for simple use databases.
package simpledb

import (
	"database/sql"
)

//DB is structure of database
type DB struct {
	Db   *sql.DB
	err  error
	rows *sql.Rows
	tx   *sql.Tx
	stmt *sql.Stmt
}

//NewDatabase - create new database
func (d *DB) NewDatabase(database, basename string) error {
	d.Db, d.err = sql.Open(database, basename)
	if d.err != nil {
		return d.err
	}
	d.err = d.Db.Ping()
	if d.err != nil {
		return d.err
	}
	return nil
}

//Query - for query from database
func (d *DB) Query(query string, args ...interface{}) ([][]interface{}, error) {
	var emptyReturn [][]interface{}
	d.rows, d.err = d.Db.Query(query, args...)
	if d.err != nil {
		return emptyReturn, d.err
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
			return emptyReturn, d.err
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
		return d.err
	}
	return nil
}

//Prepare creates a prepared statement for later queries or executions.
//Multiple queries or executions may be run concurrently from the returned statement.
// The caller must call the statement's Close method when the statement is no longer needed.
func (d *DB) Prepare(exec string) error {
	d.stmt, d.err = d.Db.Prepare(exec)

	if d.err != nil {
		return d.err
	}
	return nil
}

//TxPrepare creates a prepared statement for use within a transaction.
func (d *DB) TxPrepare(exec string) error {
	d.stmt, d.err = d.tx.Prepare(exec)

	if d.err != nil {
		return d.err
	}
	return nil
}

//StmtExec executes a query that doesn't return rows. StmtExec use in transaction.
func (d *DB) StmtExec(args ...interface{}) error {
	_, d.err = d.stmt.Exec(args...)
	if d.err != nil {
		return d.err
	}
	return nil
}

//Begin starts a transaction. The default isolation level is dependent on the driver.
func (d *DB) Begin() error {
	d.tx, d.err = d.Db.Begin()
	if d.err != nil {
		return d.err
	}
	return nil
}

//Commit commits the transaction.
func (d *DB) Commit() error {
	//defer d.stmt.Close()
	d.err = d.tx.Commit()
	if d.err != nil {
		return d.err
	}
	d.stmt.Close()
	return nil
}

//Rollback aborts the transaction.
func (d *DB) Rollback() error {
	d.err = d.tx.Rollback()
	if d.err != nil {
		return d.err
	}
	//defer d.stmt.Close()
	return nil

}

//Close - close  database
func (d *DB) Close() {
	d.Db.Close()
}
