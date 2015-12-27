package main

/*
Documents
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

    // not request now
    // just a new document object
    doc, err := db.GetDocument("test")

    if err != nil {
    }

    ret, err := doc.GetInfo()

    if err != nil {
        // maybe not document
        body := map[string]interface{}{
            "title": "This is title",
            "content": "This is content",
            "number": 100,
        }

        _, err := db.CreateDocument("test", body)
        if err != nil {
            // create error
        }

    }

    fmt.Println("first hit the cache")
    ret, err = doc.GetInfo()

    body := map[string]interface{}{
        "title": "This is title",
        "content": "This is content",
        "number": 102,
    }

    fmt.Println("update will change the Etag")
    ret, err = doc.Update(body)
    if err != nil {
        // create error
    }

    fmt.Println("get document again.")
    ret, err = doc.GetInfo()

    fmt.Println("get document again..")
    ret, err = doc.GetInfo()

    fmt.Println("get document again...")
    ret, err = doc.GetInfo()

    for key, value := range ret {
        fmt.Println(key, "=", value)
    }
}