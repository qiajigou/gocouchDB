package main

/*
Database
*/

import (
    "fmt"
    "github.com/GuoJing/gocouchDB"
)

func main(){
    DSN := "http://localhost:5984"
    client := gocouchDB.NewClientByDSN(DSN)

    // if couchdb dont't need username and password
    // remove this line

    client.SetAuth("admin", "admin")

    // not request now
    // just a new object
    db, err := client.GetDatabase("duidui")

    if err != nil {
        // handle
    }

    ret, err := db.Compact()

    if err != nil {
        // handle no_db_file or no_auth
    }

    for key, value := range ret {
        fmt.Println(key, "=", value)
    }
}