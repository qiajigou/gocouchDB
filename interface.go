package gocouchDB

import (
    "io"
    "net/http"
)

type ITransport interface {
    request(method, path string, body io.Reader, headers map[string]string) (j map[string]interface{}, err error)
    addHeader(req *http.Request, header map[string]string) (r *http.Request)
    requestList(method, path string, body io.Reader, headers map[string]string) (j []string, err error)
}

type ICache interface {
    Get(key string, defaultValue interface{})(value interface{})
    Set(key string, value interface{})(ret bool)
    Delete(key string)(ret bool)
}
