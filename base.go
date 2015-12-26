package gocouchDB

import (
    "io"
    "strings"
    "encoding/json"
    "encoding/base64"
)

type ClientBase struct {
    Headers map[string]string
    Username string
    Password string
    transport ITransport
}

type ReplicateTask struct {
    Source string
    Target string
    Proxy string
    Continuous bool
    CreateTarget bool
    Cancel bool
    DocumentIDs []string
}


// set basic auth
func (cl *ClientBase)SetAuth(username, password string) {
    cl.Username = username
    cl.Password = password
}

// called everytime a new request
func (cl *ClientBase)beforeRequest() {
    if cl.Username != "" && cl.Password != "" {
        auth := []byte(cl.Username + ":" + cl.Password)
        hashed := "Basic " + base64.StdEncoding.EncodeToString(auth)
        cl.Headers["Authorization"] = string(hashed)
    }
    cl.Headers["Content-Type"] = "application/json"
}

// clean everytime after a request
func (cl *ClientBase)afterRequest() {
    cl.ClearHeaders()
}

// basic request method for calling the transport
// and clean headers every time
// so we could do something before request
// and defer close something
func (cl *ClientBase)request(method, path string, body io.Reader)(j map[string]interface{}, err error) {
    cl.beforeRequest()

    j, err = cl.transport.request(method, path, body, cl.Headers)

    defer cl.afterRequest()

    if err != nil {
        return j, err
    }

    return j, err
}

// this is tricky beacuse couchdb _all_dbs return list
// not a map[string]interface{} :(
func (cl *ClientBase)requestList(method, path string, body io.Reader)(j []string, err error) {
    cl.beforeRequest()

    j, err = cl.transport.requestList(method, path, body, cl.Headers)

    defer cl.ClearHeaders()

    if err != nil {
        return j, err
    }

    return j, err
}

// handle map[string]interface{} to json body
func (cl *ClientBase)handleBodyData(body map[string]interface{}) (ret *strings.Reader, err error) {
    json, err := json.Marshal(body)

    if err != nil {
        return ret, err
    }

    str := strings.NewReader(string(json))

    return str, err

}

// join map[string][string] to url params
func (cl *ClientBase)joinParams(path string, params map[string]string) (url string) {

    sl := make([]string, len(params))

    i := 0

    for key, value := range params {
        tmp := key + "=" + value
        sl[i] = tmp
        i = i + 1
    }

    ps := strings.Join(sl, "&")
    p := ""

    if ps == "" {
        p = path
    } else {
        p = path + "?" + ps
    }

    return p
}

// base request method for self
func (cl *ClientBase)do(method, path string, body map[string]interface{}, params map[string]string) (j map[string]interface{}, err error) {

    if params != nil {
        path = cl.joinParams(path, params)
    }

    var ir io.Reader

    if body != nil {
        ir, err = cl.handleBodyData(body)

        if err != nil {
            return j, err
        }

    } else {
        ir = nil
    }

    return cl.request(method, path, ir)
}

func (cl *ClientBase)SetHeaders(headers map[string]string) {
    cl.Headers = headers
}

func (cl *ClientBase)ClearHeaders(){
    cl.Headers = map[string]string {}
}