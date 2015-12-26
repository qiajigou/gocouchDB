package gocouchDB

import (
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

type Transport struct {
    dsn string
    http *http.Client
}

func NewTransport(dsn string) *Transport{
    http := &http.Client{}
    return &Transport{dsn, http}
}

// request for http
// return dynamic Map
func (cl *Transport)request(method, path string, body io.Reader, headers map[string]string) (j map[string]interface{}, err error) {
    fpath := cl.dsn + "/" + path

    req, err := http.NewRequest(method, fpath, body)

    if err != nil {
        return nil, err
    }

    cl.addHeader(req, headers)

    resp, err := cl.http.Do(req)

    if err != nil {
        return nil, err
    }

    b, err := ioutil.ReadAll(resp.Body)

    var js interface{}
    err = json.Unmarshal(b, &js)

    if err != nil {
        err = CouchdbError{
            fpath, method,
            "DecodeError",
            "return json could loaded by json",
        }
    }

    defer resp.Body.Close()
    m := js.(map[string]interface{})

    couchErr := m["error"]

    if couchErr != nil {
        couchErrMessage := m["reason"]
        err = CouchdbError{
            fpath, method,
            couchErr.(string), couchErrMessage.(string),
        }
    }

    return m, err
}

// request for http
// return dynamic list
func (cl *Transport)requestList(method, path string, body io.Reader, headers map[string]string) (j []string, err error) {

    fpath := cl.dsn + "/" + path

    req, err := http.NewRequest(method, fpath, body)

    if err != nil {
        return j, err
    }

    cl.addHeader(req, headers)

    resp, err := cl.http.Do(req)

    if err != nil {
        return j, err
    }

    if resp.StatusCode == 304 {
        // If-None-modify
        err = CouchdbError{
            fpath, method,
            "None modify",
            "This document is not modified",
        }
        return j, err
    }

    if resp.StatusCode > 400 {
        err = CouchdbError{
            fpath, method,
            string(resp.StatusCode),
            "Maybe no auth?",
        }
        return j, err
    }

    b, err := ioutil.ReadAll(resp.Body)

    var js []string
    err = json.Unmarshal(b, &js)

    if err != nil {
        err = CouchdbError{
            fpath, method,
            "DecodeError",
            "return json could loaded by json",
        }
        return j, err
    }

    defer resp.Body.Close()

    return js, err
}

// add header to this transport
func (cl *Transport)addHeader(req *http.Request, header map[string]string) (r *http.Request){
    for key, value := range header {
        if key != "" && value != "" {
            req.Header.Add(key, value)
        }
    }
    return req
}
