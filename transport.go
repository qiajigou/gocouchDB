package gocouch

import (
    "fmt"
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

func (cl *Transport)request(method, path string, body io.Reader, headers map[string]string) (j map[string]interface{}, err error) {

    /*
    request a path and get result
    */

    fpath := cl.dsn + "/" + path

    req, err := http.NewRequest(method, fpath, body)

    if err != nil {
        return nil, err
    }

    fmt.Println(headers)

    cl.addHeader(req, headers)

    resp, err := cl.http.Do(req)

    if err != nil {
        return nil, err
    }

    b, err := ioutil.ReadAll(resp.Body)

    var js interface{}
    err = json.Unmarshal(b, &js)


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

func (cl *Transport)addHeader(req *http.Request, header map[string]string) (r *http.Request){
    for key, value := range header {
        if key != "" && value != "" {
            req.Header.Add(key, value)
        }
    }
    return req
}
