This project is in developing...

Yet another CouchDB GO Client

0.1 is for core API

## Quick start

    // Client is your friend
    // Use Client everywhere you want :)

    import (
        "github.com/GuoJing/gocouch"
    )

    dsn := "http://localhost:5984"
    client := gocouch.NewClientByDSN(dsn)

    db, err := client.GetDatabase("duidui")

    if err != nil {}

    doc, err := db.GetDocument("test")

    body := map[string]interface{} {
        "title" : "Haha223334444",
        "body"  : "Not at all",
    }

    doc.Update(body)

    // Or
    // doc.Delete()

## Transport

    //interface ITransport! <- you can do it yourself
    transport := NewTransport(dsn)

    client := gocouch.NewClientByTransport(transport)

## Visit without Client

    transport := NewTransport(dsn)
    db := NewDatabase(dbName, transport)
    doc, err := db.GetDocument(Name)

    doc = NewDocument(dbName, Name, transport)
    doc.Delete()

## Auth

    client := gocouch.NewClientByDSN(dsn)
    client.SetAuth(Username, Password)
    client.GetDatabase("duidui")
    // ...

    // or

    transport := NewTransport(dsn)
    db := NewDatabase(dbName, transport)
    db.SetAuth(Username, Password)
    db.GetDocument("doc")
    // ...

## TODO

1. Bulk set/get
2. Replicator
3. ETag support
4. Cache support (memory/memcached/redis)
