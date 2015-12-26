package gocouchDB

import "fmt"

type CouchdbError struct {
    path   string
    method string
    error  string
    reason string
}

func (e CouchdbError) Error() string {
    return fmt.Sprintf("method: %v path: %v error:%v reason:%v",
        e.method, e.path, e.error, e.reason)
}
