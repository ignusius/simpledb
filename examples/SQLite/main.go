package main

import (
    "fmt"

    "bitbucket.org/ignusius/simpledb"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db := new(simpledb.DB)
    // or db:=simpledb.DB{}
    err := db.NewDatabase("sqlite3", "test.db")
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
    fmt.Println(arr[0][2])

    db.Close()
}
