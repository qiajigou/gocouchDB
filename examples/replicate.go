package main

/*
Replicate database
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

    task := new(gocouchDB.ReplicateTask)
    task.Continuous = false
    task.CreateTarget = true
    task.Source = "duidui"
    task.Target = "duidui2"

    ret, err := client.Replicate(task)

    if err != nil {
        // handle repliace error
    }

    for key, value := range ret {
        fmt.Println(key, "=", value)
    }
}