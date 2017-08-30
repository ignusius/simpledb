package main

import (
    "fmt"

     "bitbucket.org/ignusius/simpledb"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db := new(simpledb.DB)

    // or db:=simpledb.DB{}
    err := db.NewDatabase("mysql", "root:password@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Exec("INSERT INTO data values (?,?,?,?,?)", 1, "test", "test", 2, 3)
    if err != nil {
        panic(err)
    }

    arr, err := db.Query("SELECT * FROM data")
    if err != nil {
        panic(err)
    }

    fmt.Println(arr)
    fmt.Println(arr[0][2])

}
