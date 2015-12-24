package gocouch

import (
    "io"
    "net/http"
)

type ITransport interface {
    request(method, path string, body io.Reader, headers map[string]string) (j map[string]interface{}, err error)
    addHeader(req *http.Request, header map[string]string) (r *http.Request)
}

type ICache interface {}
