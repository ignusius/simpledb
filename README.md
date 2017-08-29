# simpledb #

This library for simply use Databases (It works with drivers for database/sql).

## Example ##
SQlite for example
```SQL
CREATE TABLE data (
    "article" INTEGER NOT NULL,
    "title" TEXT NOT NULL,
    "note" TEXT NOT NULL DEFAULT (0),
    "sum" INTEGER NOT NULL DEFAULT (0)
, "reject" INTEGER)
```
example.go

```go
package main

import (
	"fmt"

	"bitbucket.org/ignusius/simpledb"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := new(simpledb.DB)
	// or db:=simpledb.DB{}
	err := db.NewDatabase("sqlite3", "gfz")
	if err != nil {
		panic(err)
	}

	err = db.Exec("INSERT INTO data values ($1,$2,$3,$4,$5)", 1,"test","test",2,3)
	if err !=nil{
		panic(err)
	}

	arr, err := db.Query("SELECT * FROM data")
	if err != nil {
		panic(err)
	}
	fmt.Println(arr)
	fmt.Println(simpledb.SQLstrconv(arr[0][2]))

	db.Close()
}
```
Output
```
[[1 [116 101 115 116] [116 101 115 116] 2 3]]
test
```