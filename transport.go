package gocouchDB

import (
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "strings"
)

type Transport struct {
    dsn string
    http *http.Client
    cache ICache
}

func NewTransport(dsn string) *Transport{
    http := &http.Client{}
    cache := NewMemCache()
    return &Transport{dsn, http, cache}
}

// request for http
// return dynamic Map
func (cl *Transport)request(method, path string, body io.Reader, headers map[string]string) (j map[string]interface{}, err error) {
    fpath := cl.dsn + "/" + path
    pathKey := path

    req, err := http.NewRequest(method, fpath, body)

    if err != nil {
        return nil, err
    }

    cachedETag := cl.cache.Get(pathKey, nil)

    if cachedETag != nil {
        tmp := cachedETag.(string)
        headers["If-None-Match"] = tmp
    }

    cl.addHeader(req, headers)

    resp, err := cl.http.Do(req)

    if err != nil {
        return nil, err
    }

    if resp.StatusCode == 304 {
        valueKey := cachedETag.(string)
        r := cl.cache.Get(valueKey, nil)
        if r != nil {
            return r.(map[string]interface{}), err
        }
    } else {
        if cachedETag != nil {
            valueKey := cachedETag.(string)
            cl.cache.Delete(valueKey)
        }
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

    var js interface{}
    err = json.Unmarshal(b, &js)

    if err != nil {
        err = CouchdbError{
            fpath, method,
            "DecodeError",
            "return json could loaded by json",
        }
    }

    etag := resp.Header["Etag"]

    // set cache in memory
    // path -> etag
    // etag -> json value
    if etag != nil && method == GET {
        valueKey := strings.Join(etag, "")
        cl.cache.Set(pathKey, valueKey)
        cl.cache.Set(valueKey, js)
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
        // If-None-Match
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
